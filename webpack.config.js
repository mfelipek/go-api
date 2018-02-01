const path = require('path');

module.exports = {
  entry: {
    main: path.resolve(__dirname, './react/main.js'),
  },
 output: {
    path: path.join(__dirname, 'static/scripts'),
    filename: '[name].bundle.js',
    chunkFilename: '[id].chunk.js'
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        use: {
        	loader: 'babel-loader',
            options: {
              presets: ['@babel/preset-env']
            }
        }
      }
    ]
  }
};