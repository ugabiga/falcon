import {ApolloClient, ApolloLink, HttpLink, InMemoryCache} from "@apollo/client";


const GQL_ENDPOINT = 'http://localhost:8080/graph';

function getCookie(name: string) {
    let cookieValue = "";
    if (document.cookie && document.cookie !== '') {
        const cookies = document.cookie.split(';');
        for (let i = 0; i < cookies.length; i++) {
            const cookie = cookies[i].trim();
            // Does this cookie string begin with the name we want?
            if (cookie.substring(0, name.length + 1) === (name + '=')) {
                cookieValue = decodeURIComponent(cookie.substring(name.length + 1));
                break;
            }
        }
    }
    return cookieValue;
}

const gqlClient = new HttpLink({
    uri: GQL_ENDPOINT,
    headers: {
        'Content-Type': 'application/json',
    },
    credentials: "include"
})

const cache: InMemoryCache = new InMemoryCache();

export const client = new ApolloClient({
    link: ApolloLink.split(
        operation => operation.getContext().clientName === 'gqlClient',
        gqlClient,
        gqlClient,
    ),
    cache: cache,
});