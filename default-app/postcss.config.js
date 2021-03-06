const production = (process.env.NODE_ENV == 'production')

module.exports = {
    plugins: [
        require('postcss-import')(),
        require('tailwindcss'),
        require('autoprefixer'),
        ...(production ? [require('@fullhuman/postcss-purgecss')({

            // Specify the paths to all of the template files in your project
            content: [
                './src/**/*.html',
            ],

            // Include any special characters you're using in this regular expression
            defaultExtractor: content => content.match(/[\w-/:]+(?<!:)/g) || []
        }),
        require('cssnano')({
            preset: 'default',
        })] : []),
    ],
}
