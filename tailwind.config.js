/** @type {import('tailwindcss').Config} */

module.exports = {
    content: [ "./**/*.html", "./**/*.templ", "./**/*.go", ],
   safelist: [],
   theme: {
       extend: {
          fontFamily:{
            "sans": ["Averta" ,"Roboto", "sans"],
            "serif": ["Averta" ,"Roboto", "serif"],
          },
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
       themes: [
        {
            mytheme: {
              "primary": "#B2C2C1",
              "secondary": "#f6d860",
              "accent": "#37cdbe",
              "neutral": "#3d4451",
              "base-100": "#ffffff",
            },
          },
        "light"]
   }
}