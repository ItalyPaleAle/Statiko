/*
Copyright © 2020 Alessandro Segala (@ItalyPaleAle)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, version 3 of the License.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package certificates

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"sort"
	"time"

	pb "github.com/statiko-dev/statiko/shared/proto"
)

// GetX509 returns a X509 object for a certificate
func GetX509(cert []byte) (certX509 *x509.Certificate, err error) {
	// Get the certificate's x509 object
	block, _ := pem.Decode(cert)
	if block == nil {
		err = errors.New("invalid certificate PEM block")
		return nil, err
	}
	certX509, err = x509.ParseCertificate(block.Bytes)
	return
}

// InspectCertificate loads a X.509 certificate and checks its details, such as expiration
func InspectCertificate(site *pb.Site, obj *pb.TLSCertificate, cert *x509.Certificate) error {
	now := time.Now()

	// Check "NotAfter" (require at least 12 hours)
	if cert.NotAfter.Before(now.Add(12 * time.Hour)) {
		return fmt.Errorf("certificate has expired or has less than 12 hours of validity: %v", cert.NotAfter)
	}

	// Check "NotBefore"
	if !cert.NotBefore.Before(now) {
		return fmt.Errorf("certificate's NotBefore is in the future: %v", cert.NotBefore)
	}

	// Check if the list of domains matches, but only for self-signed or ACME certificates
	// We're not checking this for imported certificates because they might have wildcards and be valid for more domains
	if obj.Type == pb.TLSCertificate_ACME || obj.Type == pb.TLSCertificate_SELF_SIGNED {
		domains := append([]string{site.Domain}, site.Aliases...)
		sort.Strings(domains)
		certDomains := append(make([]string, 0), cert.DNSNames...)
		sort.Strings(certDomains)
		if !reflect.DeepEqual(domains, certDomains) {
			return fmt.Errorf("list of domains in certificate does not match: %v", certDomains)
		}
	}

	return nil
}

// GenerateTLSCert generates a new self-signed TLS certificate (with a RSA 4096-bit key) and returns the private key and public certificate encoded as PEM
// The first domain is the primary one, used as value for the "Common Name" value too
// Each certificate is valid for 1 year
func GenerateTLSCert(domains ...string) (keyPEM []byte, certPEM []byte, err error) {
	// Ensure we have at least 1 domain
	if len(domains) < 1 {
		err = errors.New("need to specify at least one domain name")
		return
	}

	// Generate a private key
	// The main() method has already invoked rand.Seed
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return
	}

	// Build the X.509 certificate
	now := time.Now()
	tpl := x509.Certificate{}
	tpl.BasicConstraintsValid = false
	tpl.DNSNames = domains
	tpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}
	tpl.KeyUsage = x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageDataEncipherment
	tpl.IsCA = false
	tpl.NotAfter = now.Add(8760 * time.Hour) // 1 year
	tpl.NotBefore = now
	tpl.SerialNumber = big.NewInt(1)
	tpl.SignatureAlgorithm = x509.SHA256WithRSA
	tpl.Subject = pkix.Name{
		Organization: []string{SelfSignedCertificateIssuer},
		CommonName:   domains[0],
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, &tpl, &tpl, &privateKey.PublicKey, privateKey)
	if err != nil {
		return
	}

	// Encode the key in a PEM block
	buf := &bytes.Buffer{}
	err = pem.Encode(buf, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err != nil {
		return
	}
	keyPEM = buf.Bytes()

	// Encode the certificate in a PEM block
	buf = &bytes.Buffer{}
	err = pem.Encode(buf, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	})
	if err != nil {
		return
	}
	certPEM = buf.Bytes()

	return
}
