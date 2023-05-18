module.exports = {
    content: [
        "./src/**/*.{js,jsx,ts,tsx}",
    ],
    theme: {
        extend: {
            colors :{
              'azul': {
                100: '#ebecf6',
                200: '#c4c6e5',
                300: '#9da0d4',
                400: '#757ac2',
                500: '#4e54b1',
                600: '#3c418a',
                700: '#2b2f62',
                800: '#1a1c3b',
                900: '#090914',
              },
              'dourado': {
                100: '#fbf6f0',
                200: '#f3e5d1',
                300: '#ead4b2',
                400: '#e2c393',
                500: '#dab274',
                600: '#d1a155',
                700: '#c98f36',
                800: '#aa792e',
                900: '#8b6325',
              },
            },
            screens: {
              'sm': '640px',
              'md': '768px',
              'lg': '1024px',
              'xl': '1280px',
              '2xl': '1536px',
            }
          },
    },
    plugins: [],
}