/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [ "./**/*.html", "./**/*.templ", "./**/*.go", ],
   safelist: [],
   theme: {
       extend: {
           colors: {
              "branded": {
                //   "100": "#E6F6FF",
                //   "200": "#BFE8FF",
                //   "300": "#99DAFF",
                  "400": "#B2C2C1",
                //   "500": "#00A6FF",
                //   "600": "#0096E6",
                //   "700": "#00647F",
                  "800": "#586A5D",
                //   "900": "#00333A",
              },
           }
       }
   },
   plugins: [require("daisyui")],
   daisyui: {
       themes: [ "light"]
   }
}