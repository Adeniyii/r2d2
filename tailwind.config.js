module.exports = {
  content: ["./cmd/web/templates/**/*.tmpl"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
};
