import express from 'express';
import axios from 'axios';
import dotenv from 'dotenv';

dotenv.config();

const MICROSERVICES = process.env.MICROSERVICES?.split(',') || [];
const PORT = process.env.PORT || 3000;

const app = express();

app.use(express.json());

app.get('/health-check', async (req, res) => {
  const statusReports = await Promise.all(MICROSERVICES.map(async (service) => {
    try {
      const response = await axios.get(`${service}/health`);
      return { service, status: response.status === 200 ? 'UP' : 'DOWN', data: response.data };
    } catch (error) {
      return { service, status: 'DOWN', error: error.message };
    }
  }));

  res.json(statusReports);
});

app.get('/performance-metrics', async (req, res) => {
  const metricsReports = await Promise.all(MICROSERVICES.map(async (service) => {
    try {
      const response = await axios.get(`${service}/metrics`);
      return { service, metrics: response.data };
    } catch (error) {
      return { service, error: error.message };
    }
  }));

  res.json(metricsReports);
});

app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});