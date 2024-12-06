/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./templates/**/*.html", "./*.go"],
    theme: {
        fontFamily: {
            roboto: ['Roboto', 'sans-serif']
        },
        extend: {
            inset: {
                '-second': '-20%',
                '-third': '-40%',
                'second': '7.5%',
                'third': '15%'
            },
        },
    },
    plugins: [
        require('@tailwindcss/typography'),
    ],
}

