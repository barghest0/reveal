import path from 'path';
import HtmlWebpackPlugin from 'html-webpack-plugin';
import { Configuration } from 'webpack';
import 'webpack-dev-server'


const config: Configuration = {
    entry: path.resolve(__dirname, "../src/index.tsx"),
    output: {
        filename: "bundle.js",
        path: path.resolve(__dirname, "dist"),
        clean: true,
    },
    resolve: {
        extensions: [".tsx", ".ts", ".jsx", ".js"],
        modules: [path.resolve(__dirname, "src"), "node_modules"]
    },
    module: {
        rules: [
            {
                test: /\.(ts|js)x?$/,
                exclude: /node_modules/,
                use: ["babel-loader"],
            }, 
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader'],
            },
        ],
    },
    mode: 'development',
    devtool: "cheap-module-source-map",
    devServer: {
        open: true
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: "public/index.html",
        }),
    ],
    
}

export default config;