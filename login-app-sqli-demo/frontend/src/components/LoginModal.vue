<template>
  <div v-if="isVisible" class="modal-overlay" @click="closeModal">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h2>
          <i class="fas fa-sign-in-alt"></i>
          Login
        </h2>
        <button class="close-btn" @click="closeModal">
          <i class="fas fa-times"></i>
        </button>
      </div>
      
      <form @submit.prevent="handleLogin" class="login-form">
        <div class="form-group">
          <label for="username">
            <i class="fas fa-user"></i>
            Username
          </label>
          <input
            type="text"
            id="username"
            v-model="formData.username"
            required
            placeholder="Enter your username"
          />
        </div>
        
        <div class="form-group">
          <label for="password">
            <i class="fas fa-lock"></i>
            Password
          </label>
          <input
            type="password"
            id="password"
            v-model="formData.password"
            required
            placeholder="Enter your password"
          />
        </div>

        <div class="endpoint-selection">
          <label>
            <input 
              type="radio" 
              v-model="useSecureEndpoint" 
              :value="false"
            />
            <span class="vulnerability-badge vulnerable">Vulnerable Endpoint</span>
          </label>
          <label>
            <input 
              type="radio" 
              v-model="useSecureEndpoint" 
              :value="true"
            />
            <span class="vulnerability-badge secure">Secure Endpoint</span>
          </label>
        </div>
        
        <div v-if="error" class="error-message">
          <i class="fas fa-exclamation-circle"></i>
          {{ error }}
        </div>
        
        <div v-if="success" class="success-message">
          <i class="fas fa-check-circle"></i>
          {{ success }}
        </div>
        
        <div class="form-actions">
          <button type="submit" class="btn-primary" :disabled="isLoading">
            <i class="fas fa-sign-in-alt" v-if="!isLoading"></i>
            <i class="fas fa-spinner fa-spin" v-else></i>
            {{ isLoading ? 'Logging in...' : 'Login' }}
          </button>
          <button type="button" class="btn-secondary" @click="closeModal">
            Cancel
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import ApiService from '../services/api.js'

export default {
  name: 'LoginModal',
  props: {
    isVisible: {
      type: Boolean,
      default: false
    }
  },
  emits: ['close', 'login-success'],
  data() {
    return {
      formData: {
        username: '',
        password: ''
      },
      useSecureEndpoint: false,
      error: '',
      success: '',
      isLoading: false
    }
  },
  methods: {
    async handleLogin() {
      this.error = ''
      this.success = ''
      this.isLoading = true
      
      try {
        let response
        if (this.useSecureEndpoint) {
          response = await ApiService.loginSecure(
            this.formData.username,
            this.formData.password
          )
        } else {
          response = await ApiService.loginVulnerable(
            this.formData.username,
            this.formData.password
          )
        }
        
        if (response.success) {
          this.success = response.message
          this.$emit('login-success', {
            user: response.user,
            endpoint: this.useSecureEndpoint ? 'secure' : 'vulnerable'
          })
          setTimeout(() => {
            this.closeModal()
          }, 1500)
        } else {
          this.error = response.message
        }
      } catch (error) {
        this.error = error.message || 'Login failed'
      } finally {
        this.isLoading = false
      }
    },
    
    closeModal() {
      this.formData = { username: '', password: '' }
      this.error = ''
      this.success = ''
      this.isLoading = false
      this.$emit('close')
    }
  }
}
</script>