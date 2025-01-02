const path = require('path');

module.exports = {
    entry: './src/ts/test.ts',
    output: {
        path: path.resolve(__dirname, './src/js/build'),
        filename: 'bundle.js',
    },
    module: {
        rules: [
            {
                test: /\.ts$/,
                use: 'ts-loader',
                exclude: /node_modules/,
            },
        ],
    },
    resolve: {
        extensions: ['.ts', '.js'],
    },
};