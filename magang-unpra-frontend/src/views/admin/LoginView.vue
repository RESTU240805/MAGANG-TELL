<template>
  <div class="min-h-screen bg-gray-900 flex items-center justify-center">
    <div class="bg-white rounded-2xl shadow-xl p-10 w-full max-w-md">
      
      <!-- Logo -->
      <div class="text-center mb-8">
        <span class="text-3xl font-black text-green-600">TeLpp.</span>
        <p class="text-gray-500 text-sm mt-1">Admin Portal</p>
      </div>

      <!-- Form -->
      <div class="space-y-4">
        <div>
          <label class="text-sm font-medium text-gray-700 block mb-1">Email</label>
          <input v-model="form.email" type="email" placeholder="admin@telpp.com"
            class="w-full border border-gray-300 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
        </div>
        <div>
          <label class="text-sm font-medium text-gray-700 block mb-1">Password</label>
          <input v-model="form.password" type="password" placeholder="••••••••"
            class="w-full border border-gray-300 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
        </div>

        <p v-if="error" class="text-red-500 text-sm">{{ error }}</p>

        <button @click="handleLogin"
          class="w-full bg-green-600 text-white py-3 rounded-xl font-semibold hover:bg-green-700 transition mt-2">
          {{ loading ? 'Loading...' : 'Login' }}
        </button>
      </div>

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

const handleLogin = async () => {
  error.value = ''
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