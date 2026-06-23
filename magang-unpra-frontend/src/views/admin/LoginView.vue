<template>
  <div class="login-page">
    <img src="/images/gedung.jpeg" alt="" class="login-bg" />
    <div class="login-overlay"></div>

    <div class="login-card">
      <div class="login-card-header">
        <img src="/images/logo-telpp.png" alt="TELPP" class="login-logo" />
        <h2>Welcome Back</h2>
        <p>Sign in to access the admin dashboard</p>
      </div>

      <div class="login-card-body">
        <div class="login-field">
          <label>Email</label>
          <div class="login-input-wrap">
            <svg class="login-input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
              <rect x="2" y="4" width="20" height="16" rx="3"/><polyline points="22,6 12,13 2,6"/>
            </svg>
            <input v-model="form.email" type="email" placeholder="admin@telpp.com" @keyup.enter="handleLogin" />
          </div>
        </div>

        <div class="login-field">
          <label>Password</label>
          <div class="login-input-wrap">
            <svg class="login-input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
              <rect x="3" y="11" width="18" height="11" rx="2"/><path d="M7 11V7a5 5 0 0110 0v4"/>
            </svg>
            <input v-model="form.password" :type="showPass ? 'text' : 'password'" placeholder="Enter your password" @keyup.enter="handleLogin" />
            <button type="button" class="login-eye" @click="showPass = !showPass">
              <svg v-if="showPass" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" class="w-5 h-5">
                <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8S1 12 1 12z"/><circle cx="12" cy="12" r="3"/>
              </svg>
              <svg v-else viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" class="w-5 h-5">
                <path d="M17.94 17.94A10.07 10.07 0 0112 20c-7 0-11-8-11-8a18.45 18.45 0 015.06-5.94M9.9 4.24A9.12 9.12 0 0112 4c7 0 11 8 11 8a18.5 18.5 0 01-2.16 3.19m-6.72-1.07a3 3 0 11-4.24-4.24"/>
                <line x1="1" y1="1" x2="23" y2="23"/>
              </svg>
            </button>
          </div>
        </div>

        <Transition name="shake">
          <div v-if="error" class="login-error">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" class="w-4 h-4">
              <circle cx="12" cy="12" r="10"/><line x1="15" y1="9" x2="9" y2="15"/><line x1="9" y1="9" x2="15" y2="15"/>
            </svg>
            {{ error }}
          </div>
        </Transition>

        <button @click="handleLogin" :disabled="loading" class="login-btn">
          <div v-if="loading" class="login-spinner"></div>
          <span>{{ loading ? 'Signing in...' : 'Sign In' }}</span>
        </button>
      </div>

      <p class="login-footer">Copyright 2026 PT TELPP. All rights reserved.</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '../../services/api'

const router = useRouter()
const form = ref({ email: '', password: '' })
const error = ref('')
const loading = ref(false)
const showPass = ref(false)

const handleLogin = async () => {
  error.value = ''
  if (!form.value.email || !form.value.password) {
    error.value = 'Email dan password wajib diisi'
    return
  }
  loading.value = true
  try {
    const res = await api.post('/login', form.value)
    localStorage.setItem('token', res.data.token)
    localStorage.setItem('user', JSON.stringify(res.data.user))
    router.push('/admin/dashboard')
  } catch (_err) {
    error.value = 'Email atau password salah'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-page {
  position: relative;
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #111;
  padding: 20px;
}
.login-bg {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.login-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(160deg, rgba(5, 28, 14, 0.92) 0%, rgba(5, 28, 14, 0.75) 50%, rgba(5, 28, 14, 0.55) 100%);
}

.login-card {
  position: relative;
  z-index: 2;
  background: #fff;
  border-radius: 20px;
  width: 100%;
  max-width: 400px;
  padding: 40px 36px 32px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
}

.login-card-header {
  text-align: center;
  margin-bottom: 32px;
}
.login-logo {
  height: 52px;
  width: auto;
  margin-bottom: 20px;
}
.login-card-header h2 {
  font-size: 1.5rem;
  font-weight: 800;
  color: #111;
  margin: 0 0 6px;
}
.login-card-header p {
  font-size: 0.82rem;
  color: #999;
  margin: 0;
}

.login-field {
  margin-bottom: 18px;
}
.login-field label {
  display: block;
  font-size: 0.78rem;
  font-weight: 600;
  color: #444;
  margin-bottom: 6px;
}
.login-input-wrap {
  position: relative;
  display: flex;
  align-items: center;
}
.login-input-icon {
  position: absolute;
  left: 14px;
  width: 18px;
  height: 18px;
  color: #aaa;
  pointer-events: none;
}
.login-input-wrap input {
  width: 100%;
  padding: 12px 14px 12px 44px;
  border: 1.5px solid #ddd;
  border-radius: 12px;
  font-size: 0.88rem;
  color: #222;
  background: #fafafa;
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s, background 0.2s;
}
.login-input-wrap input:focus {
  border-color: #5F9E42;
  box-shadow: 0 0 0 3px rgba(95, 158, 66, 0.12);
  background: #fff;
}
.login-input-wrap input::placeholder {
  color: #bbb;
}
.login-eye {
  position: absolute;
  right: 12px;
  background: none;
  border: none;
  cursor: pointer;
  color: #aaa;
  padding: 2px;
  display: flex;
  align-items: center;
  transition: color 0.2s;
}
.login-eye:hover { color: #555; }

.login-error {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 14px;
  background: #fef2f2;
  border: 1px solid #fecaca;
  border-radius: 10px;
  color: #dc2626;
  font-size: 0.82rem;
  font-weight: 500;
  margin-bottom: 14px;
}

.login-btn {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 13px;
  background: #5F9E42;
  color: #fff;
  border: none;
  border-radius: 12px;
  font-size: 0.9rem;
  font-weight: 700;
  cursor: pointer;
  transition: background 0.2s, transform 0.15s;
  margin-top: 6px;
}
.login-btn:hover { background: #4d8536; transform: translateY(-1px); }
.login-btn:active { transform: translateY(0); }
.login-btn:disabled { opacity: 0.65; cursor: not-allowed; transform: none; }
.login-spinner {
  width: 18px; height: 18px;
  border: 2.5px solid rgba(255, 255, 255, 0.3);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}
@keyframes spin { to { transform: rotate(360deg); } }

.login-footer {
  text-align: center;
  font-size: 0.72rem;
  color: #ccc;
  margin-top: 28px;
}

.shake-enter-active { animation: shake 0.4s ease; }
.shake-leave-active { transition: opacity 0.2s; }
.shake-leave-to { opacity: 0; }
@keyframes shake {
  0%, 100% { transform: translateX(0); }
  20%, 60% { transform: translateX(-6px); }
  40%, 80% { transform: translateX(6px); }
}

@media (max-width: 480px) {
  .login-card { padding: 32px 24px 24px; }
  .login-logo { height: 44px; }
}
</style>
