/** @type {import('tailwindcss').Config}*/
const config = {
  content: ["./src/**/*.{html,js,svelte,ts}"],

  theme: {
    extend: {
      backgroundImage: {
        'tree_login': "url('/tree_login.png')",
      }
    },
    
  },

  plugins: [],
};

module.exports = config;
