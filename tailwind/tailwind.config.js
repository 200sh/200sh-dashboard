/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./templates/**/*.html", "./*.go", "./views/**/*.go"],
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
            colors: {
                primary: "#9ACD32",
        },
        }
    },
    plugins: [
        require('@tailwindcss/typography'),
    ],
}
