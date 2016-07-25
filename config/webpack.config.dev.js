const path = require('path');
// const webpack = require('webpack');
const autoprefixer = require('autoprefixer');

// const HtmlWebpackPlugin = require('html-webpack-plugin');

// App file paths
const PATHS = {
  app: path.resolve(__dirname, '../client/js'),
  styles: path.resolve(__dirname, '../client/styles'),
  build: path.resolve(__dirname, '../build-client'),
  node_modules: path.resolve(__dirname, '../node_modules')
};

// const plugins = [
  // Shared code
  // new webpack.optimize.CommonsChunkPlugin('vendor', 'js/vendor.bundle.js'),
  // // make sure we can use Promise / fetch without importing them
  // // Always use bluebird even if native promises exist, only use fetch if we need to
  // new webpack.ProvidePlugin({
  //   'Promise': 'bluebird',
  //   'fetch': 'imports?this=>global!exports?global.fetch!whatwg-fetch'
  // }),
  // // Avoid publishing files when compilation fails
  // new webpack.NoErrorsPlugin(),
//   new webpack.DefinePlugin({
//     'process.env.NODE_ENV': JSON.stringify('development'),
//     __DEV__: JSON.stringify(JSON.parse(process.env.DEBUG || 'true'))
//   }),
//   // new webpack.optimize.OccurrenceOrderPlugin(),
//   // Populate the HTML file with the css / js files
//   // new HtmlWebpackPlugin({
//   //   template: 'src/index.html'
//   // })
// ];

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
    // filename: 'js/[name].js',
    // publicPath: '/'
  },
  // stats: {
  //   colors: true,
  //   reasons: true
  // },
  resolve: {
    // // make sure everything's using the same version of React
    // alias: {
    //   'react': path.join(__dirname, '..', 'node_modules', 'react')
    // },
    // allow you to require('file') instead of require('file.jsx')
    extensions: ['', '.js', '.jsx'],
  },
  module: {
    loaders: [
      {
        test: /\.jsx?$/,
        exclude: /node_modules/,

        // loaders: ['react-hot', 'babel', 'eslint'],
        loader: 'babel',
        // include: PATHS.app
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
  // plugins: plugins,
  postcss: function() {
    return [autoprefixer({
      browsers: ['last 2 versions']
    })]
  },
  sassLoader: {
    includePaths: [
      path.join(PATHS.node_modules, 'normalize-scss', 'sass'),
      path.join(PATHS.node_modules, 'normalize-scss', 'node_modules', 'support-for', 'sass')
    ]
  },
  // devServer: {
  //   contentBase: path.resolve(__dirname, '../src'),
  //   port: 3000,
  //   historyApiFallback: true
  // },
  // devtool: 'eval'
};
