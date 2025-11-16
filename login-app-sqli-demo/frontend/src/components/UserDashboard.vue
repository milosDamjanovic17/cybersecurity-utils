<template>
  <div class="dashboard">
    <div class="dashboard-header">
      <h2>
        <i class="fas fa-tachometer-alt"></i>
        Welcome, {{ user.username }}!
      </h2>
      <button class="btn-logout" @click="logout">
        <i class="fas fa-sign-out-alt"></i>
        Logout
      </button>
    </div>
    
    <div class="user-info">
      <div class="info-card">
        <h3>User Information</h3>
        <p><strong>Username:</strong> {{ user.username }}</p>
        <p><strong>Email:</strong> {{ user.email }}</p>
        <p><strong>Login Method:</strong> 
          <span :class="endpointClass">{{ endpointText }}</span>
        </p>
        <p><strong>Member Since:</strong> {{ formatDate(user.created_at) }}</p>
      </div>
    </div>

    <div class="demo-section">
      <h3>SQL Injection Demo</h3>
      <p>This user logged in using the <strong>{{ endpointText }}</strong> endpoint.</p>
      
      <div v-if="loginEndpoint === 'vulnerable'" class="vulnerability-warning">
        <i class="fas fa-exclamation-triangle"></i>
        <strong>Security Warning:</strong> The login was processed through a vulnerable endpoint 
        that is susceptible to SQL injection attacks.
      </div>
      
      <div v-else class="security-notice">
        <i class="fas fa-shield-alt"></i>
        <strong>Secure:</strong> The login was processed through a secure endpoint 
        using prepared statements.
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'UserDashboard',
  props: {
    user: {
      type: Object,
      required: true
    },
    loginEndpoint: {
      type: String,
      required: true
    }
  },
  emits: ['logout'],
  computed: {
    endpointText() {
      return this.loginEndpoint === 'secure' ? 'Secure Endpoint' : 'Vulnerable Endpoint'
    },
    endpointClass() {
      return this.loginEndpoint === 'secure' ? 'secure-badge' : 'vulnerable-badge'
    }
  },
  methods: {
    logout() {
      this.$emit('logout')
    },
    formatDate(dateString) {
      return new Date(dateString).toLocaleDateString()
    }
  }
}
</script>