import React from 'react';
import ReactDOM from 'react-dom';
import AppRouter from './AppRouter';
import { BrowserRouter } from 'react-router-dom';

const apiBaseUrl = process.env.REACT_APP_API_BASE_URL || 'http://localhost:5000';

function MainApp() {
  return (
    <React.StrictMode>
      <BrowserRouter>
        <AppRouter apiBaseUrl={apiBaseUrl} />
      </BrowserRouter>
    </React.StrictMode>
  );
}

ReactDOM.render(<MainApp />, document.getElementById('root'));