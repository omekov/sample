const path = require('path')
// для того чтобы загрузить .env
require('dotenv').config()
// сам webpack
import webpack from 'webpack'
// Для чтения html файлов. И его настроит 
import HtmlWebpackPlugin from 'html-webpack-plugin'
// Для очиститки dist 
import { CleanWebpackPlugin } from 'clean-webpack-plugin'
// Для того чтобы создать css 
import MiniCssExtractPlugin from 'mini-css-extract-plugin'
// Для копирование. Пример с src в dist
import CopyWebpackPlugin from 'copy-webpack-plugin'
// Для минимизации css в одну строку
import OptimizeCssAssetWebpackPlugin from 'optimize-css-assets-webpack-plugin'
// для минимизации js в одну строку
import TerserWebpackPlugin from 'terser-webpack-plugin'
// для просмотра оптимизаций
import { BundleAnalyzerPlugin } from 'webpack-bundle-analyzer'
// Вместо eslint-loader
import ESLintPlugin, {Options} from 'eslint-webpack-plugin'

// для очистки лишних css
import CleanCSS from 'clean-css'
const PORT = parseInt(process.env.PORT || '8080')
const mode = process.env.NODE_ENV == 'production' ? 'production' : 'development'
const HOST = process.env.HOST || 'localhost'
const isDev = mode === 'development'
const isPord = mode === 'production'
// Чтобы разделить по чанком, если встречается lazy load
const optimization = (): webpack.Options.Optimization => {
    return {
        splitChunks: {
            chunks: 'all'
        },
        minimizer: isPord ? [
            new OptimizeCssAssetWebpackPlugin({
                // assetNameRegExp: /\.css$/g,
                // cssProcessor: CleanCSS,
                // cssProcessorOptions: {
                //     sourceMap: true,
                // },
                // canPrint: true,
            }),
            new TerserWebpackPlugin(),
        ] : []
    }
}
// Добавления паттерн хэша если в прод. Чтобы не кэшировались файлы
const filename = (ext = '[ext]') => isDev ? `[name].${ext}` : `[name].[hash].${ext}`
// path чтобы не повторять
const pathJoin = (path1: string, path2 = '') => path.join(__dirname, path1, path2)
// Абсолютный путь
const pathResolve = (path1: string, path2 = '') => path.resolve(__dirname, path1, path2)
// Плагины, для просмотра оптимизаций
const plugins = () => {
    const base = [
        new CleanWebpackPlugin(),
        new webpack.ProgressPlugin(),
        new HtmlWebpackPlugin({
            // Для указание параметр входа
            template: pathJoin('src', 'index.html'),
            // для минимизации html в одну строку
            minify: {
                collapseWhitespace: isPord,
                removeComments: true,
                removeRedundantAttributes: true,
                useShortDoctype: true
            }
        }),
        // чтобы собрать в один файл вендора, и разделить по чанком @import
        new MiniCssExtractPlugin({
            filename: filename('css'),
            chunkFilename: filename('css'),
        }),
        // Переместить favicon в dist
        new CopyWebpackPlugin({
            patterns: [
                {
                    from: pathResolve('src/favicon.ico'),
                    to: pathResolve('dist')
                },
                // {
                //     from: 'src/assets/**/*',
                //     to: './assets',
                //     transformPath(targetPath, absolutePath) {
                //         return targetPath.replace('src/assets', '');
                //     }
                // }
            ]
        }),
    ]
    if (isDev) {
        const options: Options = {
            extensions: ['ts', 'tsx'],
            failOnError: true,
            
        }
        base.push(new ESLintPlugin(options))
    }
    if (isPord) {
        base.push(new BundleAnalyzerPlugin())
    }
    return base
}



const config: webpack.Configuration = {
    // корень проект, откуда начинается файлы
    context: pathJoin('src'),
    // prod или dev
    mode: mode,
    entry: {
        // точка сброки проекта
        app: pathJoin('src', 'index.tsx')
    },
    // Просмотривать исходный код без компилятора 
    devtool: isDev ? 'source-map' : false,
    devServer: {
        // Точка слежки
        contentBase: pathResolve('src'),
        // управление log
        stats: {
            children: false,
            maxModules: 0
        },
        // Включить отслежку
        watchContentBase: true,
        // Порт
        port: PORT,
        // Сохранять dist в ОЗУ
        hot: isDev,
        // Открывать браузер с проектом
        open: true,
        // Отображать ошибку на фоне
        overlay: true,
        // Сохранять историяю переходов состраниц
        historyApiFallback: true,
        // 
        compress: true,
        // адрес проекта
        host: HOST
    },
    // Для указания что за проект
    target: 'web',
    // Для оптимизаций выводных файлов
    optimization: optimization(),
    // для обратки пути файлов с корни проекта
    resolve: {
        alias: {
            '../../theme.config': pathJoin('src/semantic-ui/theme.config'),
            '@components': pathJoin('src/components'),
            '@pages': pathJoin('src/pages'),
            'src/semantic-ui/site': pathJoin('src/semantic-ui/site')
        },
        // возможность не указывать форматы файлов в import
        extensions: ['.ts', '.tsx', '.js']
    },
    // лоудеры
    module: {
        rules: [
            {
                // Обработка tsx  файлов
                test: /\.tsx?$/,
                exclude: /node_modules/,
                use: 'ts-loader'
            },
            {
                // обработчка css less файлов
                test: /\.(c|le)ss$/,
                use: [
                    {
                        // Собрать в одну кучу
                        loader: MiniCssExtractPlugin.loader,
                        options: {
                            hmr: isDev,
                            reloadAll: true
                        }
                    },
                    'css-loader',
                    'less-loader',
                ],
            },
            // {
            //     test: /\.(woff|woff2|ttf|eot|svg|png|j?g|gif|ico)?$/,
            //     use: [
            //         {
            //             loader: 'file-loader',
            //             options: {
            //                 name: filename(),
            //                 publicPath: 'assets'
            //             },
            //         }
            //     ],
            // },
            {
                // возможность обращения в b64
                test: /\.(woff|woff2|ttf|eot|svg|png|j?g|gif|ico)?$/,
                use: [
                    {
                        loader: 'url-loader',
                        options: {
                            name: filename(),
                            // Если файл больше 8 кб то его не надо base64
                            limit: 8192,
                        }
                    },
                ],
            },
        ],
    },
    output: {
        // точка вывода
        filename: filename('js'),
        // указать папку
        path: pathResolve('dist')
    },
    // плагины
    plugins: plugins()
}

export default config