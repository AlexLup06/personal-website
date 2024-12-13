const path = require('path');

module.exports = {
    entry: './src/js/test.js',
    output: {
        path: path.resolve(__dirname, './src/js/build'),
        filename: 'bundle.js',
    },
    mode: "development",
};