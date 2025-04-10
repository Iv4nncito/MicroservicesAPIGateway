import express from 'express';
import axios from 'axios';
import dotenv from 'dotenv';

dotenv.config();

const MICROSERVICES = process.env.MICROSERVICES?.split(',') || [];
const PORT = process.env.PORT || 3000;

const app = express();

app.use(express.json());

async function checkServiceHealth(service: string) {
  try {
    const response = await axios.get(`${service}/health`);
    return { 
      service, 
      status: response.status === 200 ? 'UP' : 'DOWN', 
      data: response.data 
    };
  } catch (error) {
    console.error(`Error fetching health status from ${service}: ${error.message}`);
    return { 
      service, 
      status: 'DOWN', 
      error: error.message 
    };
  }
}

async function fetchServiceMetrics(service: string) {
  try {
    const response = await axios.get(`${service}/metrics`);
    return { 
      service, 
      metrics: response.data 
    };
  } catch (error) {
    console.error(`Error fetching metrics from ${service}: ${error.message}`);
    return { 
      service, 
      error: error.message 
    };
  }
}

app.get('/health-check', async (req, res) => {
  try {
    const statusReports = await Promise.all(MICROSERVICES.map(service => checkServiceHealth(service)));
    res.json(statusReports);
  } catch (error) {
    console.error(`Unexpected error in /health-check: ${error.message}`);
    res.status(500).send('An unexpected error occurred');
  }
});

app.get('/performance-metrics', async (req, res) => {
  try {
    const metricsReports = await Promise.all(MICROSERVICES.map(service => fetchServiceMetrics(service)));
    res.json(metricsReports);
  } catch (error) {
    console.error(`Unexpected error in /performance-metrics: ${error.message}`);
    res.status(500).send('An unexpected error occurred');
  }
});

app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});