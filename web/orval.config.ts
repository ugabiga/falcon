import {defineConfig} from 'orval';


export default defineConfig({
    rest: {
        output: {
            mode: 'split',
            target: './src/api/endpoints/transformer.ts',
            schemas: './src/api/model',
            client: 'react-query',
            override: {
                mutator: {
                    path: './src/api/mutator/custom-instance.ts',
                    name: 'customInstance',
                },
                query: {
                    useQuery: true,
                    useSuspenseQuery: true,
                    useSuspenseInfiniteQuery: true,
                    // useInfinite: true,
                    // useInfiniteQueryParam: 'limit',
                    options: {
                        retry: 0,
                        staleTime: 10000,
                    },
                },
            }
        },
        input: {
            target: '../docs/swagger.yaml',
        },
    },
});