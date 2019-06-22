/*
Copyright © 2019 Alessandro Segala (@ItalyPaleAle)

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

/* eslint-env mocha  */

'use strict'

const assert = require('assert')
const request = require('supertest')
const validator = require('validator')

const utils = require('../shared/utils')
const sitesData = require('../shared/sites-data')
const sharedTests = require('../shared/shared-tests')

// Read URLs from env vars
const nodeUrl = process.env.NODE_URL || 'localhost:3000'
//const nginxUrl = process.env.NGINX_URL || 'localhost'

// Auth header
const auth = 'hello world'

// Supertest instances
const nodeRequest = request('https://' + nodeUrl)
//const nginxRequest = request('https://' + nginxUrl)

// Site ids
const siteIds = {}

// Configured sites and apps
const sites = {}

// Check that the platform has been started correctly
describe('SMPlatform node', function() {
    it('Adopt node', async function() {
        // This operation can take some time
        this.timeout(10 * 1000)
        this.slow(5 * 1000)

        const response = await nodeRequest
            .post('/adopt')
            .set('Authorization', auth)
            .expect('Content-Type', /json/)
            .expect(200)

        assert(response.body)
        assert.deepStrictEqual(response.body, {message: 'Adopted'})
    })

    it('Check platform data directory', sharedTests.tests.checkDataDirectory(sites))

    it('Check platform config directory', sharedTests.tests.checkConfigDirectory())

    it('Check nginx configuration', sharedTests.tests.checkNginxConfig())

    it('Create site 1', async function() {
        // This operation can take some time
        this.timeout(30 * 1000)
        this.slow(15 * 1000)

        const response = await nodeRequest
            .post('/site')
            .set('Authorization', auth)
            .send(sitesData.site1)
            .expect('Content-Type', /json/)
            .expect(200)

        assert(response.body)
        assert.deepStrictEqual(Object.keys(response.body).sort(), ['ID', 'createdAt', 'updatedAt', 'clientCaching', 'tlsCertificate', 'domain', 'aliases'].sort()) 
        assert(validator.isUUID(response.body.ID))
        assert(validator.isISO8601(response.body.createdAt, {strict: true}))
        assert(validator.isISO8601(response.body.updatedAt, {strict: true}))
        assert.strictEqual(response.body.clientCaching, sitesData.site1.clientCaching)
        assert.strictEqual(response.body.tlsCertificate, sitesData.site1.tlsCertificate)
        assert.strictEqual(response.body.domain, sitesData.site1.domain)
        assert.deepStrictEqual(response.body.aliases.sort(), sitesData.site1.aliases.sort())

        // Store site
        siteIds.site1 = response.body.ID
        sites[response.body.ID] = response.body

        // Wait a few moments for the server to finish restarting
        await utils.waitPromise(1500)

        // Check the data directory
        await sharedTests.checkDataDirectory(sites)

        // Check the Nginx configuration
        await sharedTests.checkNginxConfig(sites)
    })

    it('Nginx is up', sharedTests.tests.checkNginxStatus())

    it('Site 1 is up', sharedTests.tests.checkNginxSite(sitesData.site1))

    it('Create site 2', async function() {
        // This operation can take some time
        this.timeout(30 * 1000)
        this.slow(15 * 1000)

        const response = await nodeRequest
            .post('/site')
            .set('Authorization', auth)
            .send(sitesData.site2)
            .expect('Content-Type', /json/)
            .expect(200)

        assert(response.body)
        assert.deepStrictEqual(Object.keys(response.body).sort(), ['ID', 'createdAt', 'updatedAt', 'clientCaching', 'tlsCertificate', 'domain', 'aliases'].sort()) 
        assert(validator.isUUID(response.body.ID))
        assert(validator.isISO8601(response.body.createdAt, {strict: true}))
        assert(validator.isISO8601(response.body.updatedAt, {strict: true}))
        assert.strictEqual(response.body.clientCaching, sitesData.site2.clientCaching)
        assert.strictEqual(response.body.tlsCertificate, sitesData.site2.tlsCertificate)
        assert.strictEqual(response.body.domain, sitesData.site2.domain)
        assert.deepStrictEqual(response.body.aliases.sort(), sitesData.site2.aliases.sort())

        // Store site
        siteIds.site2 = response.body.ID
        sites[response.body.ID] = response.body

        // Wait a few moments for the server to finish restarting
        await utils.waitPromise(1500)

        // Check the data directory
        await sharedTests.checkDataDirectory(sites)

        // Check the Nginx configuration
        await sharedTests.checkNginxConfig(sites)
    })

    it('Nginx is up', sharedTests.tests.checkNginxStatus())

    it('Site 2 is up', sharedTests.tests.checkNginxSite(sitesData.site2))
})
