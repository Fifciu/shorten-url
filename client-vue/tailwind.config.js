const buildValuesForUnit = (base, unit, amount) => {
  const obj = {};
  for (let i = 0; i <= amount; i++) {
    obj[i+''] = `${base * i}${unit}`;
  }
  return obj
};

/** @type {import('tailwindcss').Config} */
export default {
  content: [],
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    spacing: buildValuesForUnit(8, 'px', 30),
    screens: {
      'md': '550px',
      'lg': '1100px',
      'xl': '1500px',
    },
    extend: {},
  },
  plugins: [],
}

