export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx,vue}'],
  theme: {
    extend: {
      screens: {
        betterhover: { raw: '(hover: hover)' },
      },
      colors: {
        primary: '#1A2238', // Dunkles Blaugrau (vermutlich Ihr Primärfarbton)
        secondary: '#9DAAF2', // Helles Blaugrau (vermutlich Ihr Sekundärfarbton)
        accent: '#FF6A3D', // Orange (akzentfarbe)
        highlight: '#F4DB7D', // Hellgelb (hervorhebungsfarbe)
        success: '#4caf50', // Farbe für Erfolg
        fail: '#f44336', // Farbe für Fehler
        info: '#2196f3', // Farbe für Info
      },
      fontFamily: {
        regular: 'Roboto' /* für font-regular */,
        medium: 'Roboto' /* für font-medium */,
        bold: 'Roboto' /* für font-bold */,
      },
    },
  },
  plugins: [],
};
