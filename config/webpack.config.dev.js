const path = require('path');
const autoprefixer = require('autoprefixer');

// App file paths
const PATHS = {
  app: path.resolve(__dirname, '../client/js'),
  styles: path.resolve(__dirname, '../client/styles'),
  build: path.resolve(__dirname, '../build-client'),
  node_modules: path.resolve(__dirname, '../node_modules'),
};

const sassLoaders = [
  'style-loader',
  'css-loader?sourceMap',
  'postcss-loader',
  'sass-loader?outputStyle=expanded',
];

module.exports = {
  // env: process.env.NODE_ENV,
  entry: path.resolve(PATHS.app, 'index.js'),
  output: {
    path: PATHS.build,
    filename: 'bundle.js',
  },
  resolve: {
    extensions: ['', '.js', '.jsx'],
  },
  module: {
    loaders: [
      {
        test: /\.jsx?$/,
        exclude: /node_modules/,

        // loaders: ['react-hot', 'babel', 'eslint'],
        loader: 'babel-loader',
        query: {
          presets: ['es2015'],
          plugins: ['transform-runtime'],
        },
      },
      {
        test: /\.scss$/,
        loader: sassLoaders.join('!'),
      },
      {
        test: /\.css$/,
        loader: 'style-loader!css-loader!postcss-loader',
      },
      // Inline base64 for <8k images, direct URLs for the rest
      // {
      //   test: /\.(png|jpg|jpeg|gif|svg|woff|woff2)$/,
      //   loader: 'url-loader?limit=8192'
      // }
    ],
  },
  postcss() {
    return [autoprefixer({
      browsers: ['last 2 versions'],
    })];
  },
  sassLoader: {
    includePaths: [
      path.join(PATHS.node_modules, 'normalize-scss', 'sass'),
      path.join(PATHS.node_modules, 'normalize-scss', 'node_modules', 'support-for', 'sass'),
    ],
  },
  // devServer: {
  //   contentBase: path.resolve(__dirname, '../src'),
  //   port: 3000,
  //   historyApiFallback: true
  // },
  // devtool: 'eval'
};
