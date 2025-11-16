<template>
  <div id="app">
    <div v-if="!currentUser" class="landing-page">
      <header class="app-header">
        <h1>
          <i class="fas fa-shield-alt"></i>
          SQL Injection Demo App
        </h1>
        <p>A demonstration of vulnerable vs secure authentication endpoints</p>
      </header>
      
      <div class="action-buttons">
        <button class="btn-primary large" @click="showLoginModal">
          <i class="fas fa-sign-in-alt"></i>
          Login
        </button>
        <button class="btn-secondary large" @click="showRegisterModal">
          <i class="fas fa-user-plus"></i>
          Register
        </button>
      </div>
      
      <div class="demo-info">
        <div class="info-card">
          <h3><i class="fas fa-info-circle"></i> Demo Purpose</h3>
          <p>This application demonstrates the difference between vulnerable and secure authentication endpoints:</p>
          <ul>
            <li><span class="vulnerable-badge">Vulnerable</span> - Susceptible to SQL injection attacks</li>
            <li><span class="secure-badge">Secure</span> - Protected with prepared statements</li>
          </ul>
        </div>
      </div>
    </div>

    <!-- User Dashboard -->
    <UserDashboard
      v-if="currentUser"
      :user="currentUser"
      :login-endpoint="loginEndpoint"
      @logout="handleLogout"
    />

    <!-- Modals -->
    <LoginModal
      :is-visible="isLoginModalVisible"
      @close="hideLoginModal"
      @login-success="handleLoginSuccess"
    />
    
    <RegisterModal
      :is-visible="isRegisterModalVisible"
      @close="hideRegisterModal"
      @register-success="handleRegisterSuccess"
    />
  </div>
</template>

<script>
import LoginModal from './components/LoginModal.vue'
import RegisterModal from './components/RegisterModal.vue'
import UserDashboard from './components/UserDashboard.vue'

export default {
  name: 'App',
  components: {
    LoginModal,
    RegisterModal,
    UserDashboard
  },
  data() {
    return {
      isLoginModalVisible: false,
      isRegisterModalVisible: false,
      currentUser: null,
      loginEndpoint: ''
    }
  },
  methods: {
    showLoginModal() {
      this.isLoginModalVisible = true
    },
    hideLoginModal() {
      this.isLoginModalVisible = false
    },
    showRegisterModal() {
      this.isRegisterModalVisible = true
    },
    hideRegisterModal() {
      this.isRegisterModalVisible = false
    },
    handleLoginSuccess(data) {
      this.currentUser = data.user
      this.loginEndpoint = data.endpoint
      this.hideLoginModal()
    },
    handleRegisterSuccess(data) {
      this.hideRegisterModal()
      // Could automatically show login modal or success message
      alert(`Registration successful using ${data.endpoint} endpoint!`)
    },
    handleLogout() {
      this.currentUser = null
      this.loginEndpoint = ''
    }
  }
}
</script>

<style>
/* Global Styles */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
}

#app {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

/* Landing Page */
.landing-page {
  max-width: 600px;
  width: 100%;
  text-align: center;
}

.app-header h1 {
  color: white;
  font-size: 2.5em;
  margin-bottom: 10px;
  text-shadow: 0 2px 4px rgba(0,0,0,0.3);
}

.app-header p {
  color: rgba(255,255,255,0.9);
  font-size: 1.1em;
  margin-bottom: 40px;
}

.action-buttons {
  display: flex;
  gap: 20px;
  margin-bottom: 40px;
  justify-content: center;
}

.btn-primary, .btn-secondary {
  padding: 12px 24px;
  border: none;
  border-radius: 6px;
  font-size: 1em;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 8px;
}

.btn-primary.large, .btn-secondary.large {
  padding: 16px 32px;
  font-size: 1.1em;
}

.btn-primary {
  background: #4CAF50;
  color: white;
}

.btn-primary:hover {
  background: #45a049;
  transform: translateY(-2px);
}

.btn-secondary {
  background: #2196F3;
  color: white;
}

.btn-secondary:hover {
  background: #1976D2;
  transform: translateY(-2px);
}

.btn-logout {
  background: #f44336;
  color: white;
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
}

/* Modal Styles */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0,0,0,0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  padding: 0;
  border-radius: 8px;
  max-width: 400px;
  width: 90%;
  max-height: 90vh;
  overflow-y: auto;
  box-shadow: 0 4px 20px rgba(0,0,0,0.3);
}

.modal-header {
  background: #f5f5f5;
  padding: 20px;
  border-bottom: 1px solid #eee;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-header h2 {
  margin: 0;
  display: flex;
  align-items: center;
  gap: 10px;
  color: #333;
}

.close-btn {
  background: none;
  border: none;
  font-size: 1.2em;
  cursor: pointer;
  color: #999;
}

.close-btn:hover {
  color: #333;
}

/* Form Styles */
.login-form, .register-form {
  padding: 20px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
  color: #333;
  display: flex;
  align-items: center;
  gap: 8px;
}

.form-group input {
  width: 100%;
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1em;
}

.form-group input:focus {
  outline: none;
  border-color: #4CAF50;
}

.endpoint-selection {
  margin: 20px 0;
  padding: 15px;
  background: #f9f9f9;
  border-radius: 4px;
}

.endpoint-selection label {
  display: block;
  margin-bottom: 8px;
  cursor: pointer;
}

.endpoint-selection input[type="radio"] {
  width: auto;
  margin-right: 8px;
}

.vulnerability-badge {
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 0.9em;
  font-weight: bold;
}

.vulnerability-badge.vulnerable {
  background: #ffebee;
  color: #c62828;
  border: 1px solid #ef9a9a;
}

.vulnerability-badge.secure {
  background: #e8f5e8;
  color: #2e7d32;
  border: 1px solid #a5d6a7;
}

.vulnerable-badge {
  background: #ffebee;
  color: #c62828;
  padding: 2px 6px;
  border-radius: 3px;
  font-size: 0.9em;
  font-weight: bold;
}

.secure-badge {
  background: #e8f5e8;
  color: #2e7d32;
  padding: 2px 6px;
  border-radius: 3px;
  font-size: 0.9em;
  font-weight: bold;
}

.error-message, .success-message {
  padding: 10px;
  border-radius: 4px;
  margin-bottom: 15px;
  display: flex;
  align-items: center;
  gap: 8px;
}

.error-message {
  background: #ffebee;
  color: #c62828;
  border: 1px solid #ef9a9a;
}

.success-message {
  background: #e8f5e8;
  color: #2e7d32;
  border: 1px solid #a5d6a7;
}

.form-actions {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
  margin-top: 20px;
}

/* Info Cards */
.info-card, .demo-info .info-card {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
}

.demo-info .info-card h3 {
  margin-bottom: 15px;
  color: #333;
  display: flex;
  align-items: center;
  gap: 8px;
}

.demo-info .info-card ul {
  margin-left: 20px;
}

.demo-info .info-card li {
  margin-bottom: 8px;
}

/* Dashboard Styles */
.dashboard {
  max-width: 800px;
  width: 100%;
  background: white;
  border-radius: 8px;
  padding: 0;
  box-shadow: 0 4px 20px rgba(0,0,0,0.1);
}

.dashboard-header {
  background: #f5f5f5;
  padding: 20px;
  border-bottom: 1px solid #eee;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-radius: 8px 8px 0 0;
}

.dashboard-header h2 {
  margin: 0;
  color: #333;
  display: flex;
  align-items: center;
  gap: 10px;
}

.user-info, .demo-section {
  padding: 20px;
}

.demo-section {
  border-top: 1px solid #eee;
}

.vulnerability-warning {
  padding: 15px;
  background: #fff3cd;
  border: 1px solid #ffeaa7;
  border-radius: 4px;
  color: #856404;
  margin-top: 15px;
  display: flex;
  align-items: flex-start;
  gap: 10px;
}

.security-notice {
  padding: 15px;
  background: #d4edda;
  border: 1px solid #c3e6cb;
  border-radius: 4px;
  color: #155724;
  margin-top: 15px;
  display: flex;
  align-items: flex-start;
  gap: 10px;
}

/* Responsive */
@media (max-width: 600px) {
  .action-buttons {
    flex-direction: column;
  }
  
  .dashboard-header {
    flex-direction: column;
    gap: 15px;
    text-align: center;
  }
}
</style>