import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import './index.css';

const getEnvVariable = (key, defaultValue = '') => {
  return process.env[key] || defaultValue;
};

const API_BASE_URL = getEnvVariable('REACT_APP_API_BASE_URL');

console.log(`API Base URL: ${API_BASE_URL}`);

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);