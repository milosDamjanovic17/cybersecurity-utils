<template>
  <div v-if="isVisible" class="modal-overlay" @click="closeModal">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h2>
          <i class="fas fa-user-plus"></i>
          Register
        </h2>
        <button class="close-btn" @click="closeModal">
          <i class="fas fa-times"></i>
        </button>
      </div>
      
      <form @submit.prevent="handleRegister" class="register-form">
        <div class="form-group">
          <label for="reg-username">
            <i class="fas fa-user"></i>
            Username
          </label>
          <input
            type="text"
            id="reg-username"
            v-model="formData.username"
            required
            placeholder="Enter username"
          />
        </div>
        
        <div class="form-group">
          <label for="reg-email">
            <i class="fas fa-envelope"></i>
            Email
          </label>
          <input
            type="email"
            id="reg-email"
            v-model="formData.email"
            required
            placeholder="Enter your email"
          />
        </div>
        
        <div class="form-group">
          <label for="reg-password">
            <i class="fas fa-lock"></i>
            Password
          </label>
          <input
            type="password"
            id="reg-password"
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
            <i class="fas fa-user-plus" v-if="!isLoading"></i>
            <i class="fas fa-spinner fa-spin" v-else></i>
            {{ isLoading ? 'Registering...' : 'Register' }}
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
  name: 'RegisterModal',
  props: {
    isVisible: {
      type: Boolean,
      default: false
    }
  },
  emits: ['close', 'register-success'],
  data() {
    return {
      formData: {
        username: '',
        email: '',
        password: ''
      },
      useSecureEndpoint: false,
      error: '',
      success: '',
      isLoading: false
    }
  },
  methods: {
    async handleRegister() {
      this.error = ''
      this.success = ''
      this.isLoading = true
      
      try {
        let response
        if (this.useSecureEndpoint) {
          response = await ApiService.registerSecure(
            this.formData.username,
            this.formData.email,
            this.formData.password
          )
        } else {
          response = await ApiService.registerVulnerable(
            this.formData.username,
            this.formData.email,
            this.formData.password
          )
        }
        
        if (response.success) {
          this.success = response.message
          this.$emit('register-success', {
            endpoint: this.useSecureEndpoint ? 'secure' : 'vulnerable'
          })
          setTimeout(() => {
            this.closeModal()
          }, 1500)
        } else {
          this.error = response.message
        }
      } catch (error) {
        this.error = error.message || 'Registration failed'
      } finally {
        this.isLoading = false
      }
    },
    
    closeModal() {
      this.formData = { username: '', email: '', password: '' }
      this.error = ''
      this.success = ''
      this.isLoading = false
      this.$emit('close')
    }
  }
}
</script>