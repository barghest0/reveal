import React from "react";
import ReactDOM from 'react-dom/client';
import { App } from "./app/app";
import { Provider } from "react-redux";
import { store } from "app/store";

const container = document.getElementById('root');

const root = ReactDOM.createRoot(container as HTMLDivElement);

root.render(
    <React.StrictMode>
        <Provider store={store}>
            <App/>
        </Provider>
    </React.StrictMode>
);

