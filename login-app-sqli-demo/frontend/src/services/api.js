import axios from 'axios'

const API_BASE_URL = 'http://localhost:8080/api'

class ApiService {
  constructor() {
    this.client = axios.create({
      baseURL: API_BASE_URL,
      headers: {
        'Content-Type': 'application/json'
      }
    })
  }

  // Vulnerable endpoints
  async loginVulnerable(username, password) {
    try {
      const response = await this.client.post('/login-vulnerable', {
        username,
        password
      })
      return response.data
    } catch (error) {
      throw error.response?.data || { success: false, message: 'Network error' }
    }
  }

  async registerVulnerable(username, email, password) {
    try {
      const response = await this.client.post('/register-vulnerable', {
        username,
        email,
        password
      })
      return response.data
    } catch (error) {
      throw error.response?.data || { success: false, message: 'Network error' }
    }
  }

  // Secure endpoints
  async loginSecure(username, password) {
    try {
      const response = await this.client.post('/login-secure', {
        username,
        password
      })
      return response.data
    } catch (error) {
      throw error.response?.data || { success: false, message: 'Network error' }
    }
  }

  async registerSecure(username, email, password) {
    try {
      const response = await this.client.post('/register-secure', {
        username,
        email,
        password
      })
      return response.data
    } catch (error) {
      throw error.response?.data || { success: false, message: 'Network error' }
    }
  }

  // Get users (for demonstration)
  async getUsersVulnerable(username = '') {
    try {
      const url = username ? `/users-vulnerable?username=${encodeURIComponent(username)}` : '/users-vulnerable'
      const response = await this.client.get(url)
      return response.data
    } catch (error) {
      throw error.response?.data || { success: false, message: 'Network error' }
    }
  }
}

export default new ApiService()