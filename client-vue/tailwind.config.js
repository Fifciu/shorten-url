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
      'md': `${550 / 16}rem`,
      'lg': `${1100 / 16}rem`,
      'xl': `${1500 / 16}rem`,
    },
    colors: {
      blue: {
        50: '#F4F6FA',
        500: '#0043C3'
      },
      orange: '#EB5D27',
      gray: {
        50: '#F5F5F5',
        100: '#F2F0F0',
        200: '#C7C4C4',
        400: '#8F8B8B',
        500: '#5A5757',
        750: '#1C1C1C',
      },
      white: '#FFF'
    }
    // extend: {},
  },
  plugins: [],
}

