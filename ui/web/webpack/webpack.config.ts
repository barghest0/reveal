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
        modules: [path.resolve(__dirname, "src"), "node_modules"],
        alias: {
            app: path.resolve(__dirname, "../src/app/"),
            entities: path.resolve(__dirname, "../src/entities/"),
            features: path.resolve(__dirname, "../src/features/"),
            pages: path.resolve(__dirname, "../src/pages/"),
            shared: path.resolve(__dirname, "../src/shared/"),
            widgets: path.resolve(__dirname, "../src/widgets/"),
        }
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
            {
                test: /\.(woff2?|eot|ttf|otf)$/i,
                type: 'asset/resource',
                generator: {
                    filename: 'assets/fonts/[name][ext]',
                }
            }
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