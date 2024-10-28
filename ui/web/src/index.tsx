import React from "react";
import ReactDOM from 'react-dom';
import { App } from "./app";

const container = document.getElementById('root');

const root = ReactDOM.createRoot(container as HTMLDivElement);

root.render(
    <React.StrictMode>
        <App/>
    </React.StrictMode>
);