import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://localhost:8080', // Update with your API base URL
});

export default instance;