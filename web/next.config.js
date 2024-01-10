const path = require("path");
const withPWAInit = require("next-pwa");
const {i18n} = require('./next-i18next.config')

/** @type {import('next-pwa').PWAConfig} */
const withPWA = withPWAInit({
    dest: "public",
    // Solution: https://github.com/shadowwalker/next-pwa/issues/424#issuecomment-1399683017
    register: true,
    skipWaiting: true,
    buildExcludes: ["app-build-manifest.json"],
    disable: process.env.NODE_ENV === "development",
});


/** @type {import('next').NextConfig} */
const nextConfig = {
    // i18n
    webpack: (config, {isServer}) => {
        config.externals.push({
            'utf-8-validate': 'commonjs utf-8-validate',
            'bufferutil': 'commonjs bufferutil',
        })
        if (!isServer) {
            config.resolve.fallback.fs = false
            config.resolve.fallback.net = false
            config.resolve.fallback.tls = false
            config.resolve.fallback.dns = false
        }
        return config
    },
    async headers() {
        return [
            {
                source: '/:path*',
                headers: [
                    {
                        // https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP
                        key: 'Content-Security-Policy',
                        value: "default-src * 'self' data: 'unsafe-inline' 'unsafe-eval' *"
                    }
                ]
            }
        ]
    }
}

module.exports = withPWA(nextConfig);