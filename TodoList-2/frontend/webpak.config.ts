const path = require("path");

module.exports = {
  mode: "development",
  entry: "frontend/src/index.tsx",
  output: {
    path: path.resolve(__dirname, "dist"),
    filename: "bundle.js",
  },

  resolve: {
    extensions: [".ts", ".tsx", ".js", ".jsx"],
  },
  module: {
    rules: [
      {
        test: /\.(ts|tsx)$/,
        exclude: /node_modules/,
        use: {
          loader: "ts-loader",
          options: {
            transpileOnly: true, // Set to true if you have other loaders that handle type checking
          },
        },
      },
    ],
  },
};
