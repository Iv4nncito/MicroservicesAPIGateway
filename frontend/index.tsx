import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import './index.css';

const API_BASE_URL = process.env.REACT_APP_API_BASE_URL;

console.log(`API Base URL: ${API_BASE_URL}`);

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);