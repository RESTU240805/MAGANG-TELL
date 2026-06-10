<template>
  <div class="min-h-screen bg-gray-100 flex">

    <!-- Sidebar -->
    <aside class="w-64 bg-gray-950 text-white flex flex-col fixed h-full">
      <div class="p-6 border-b border-gray-800">
        <div class="flex items-center gap-3">
          <div class="w-9 h-9 bg-green-600 rounded-xl flex items-center justify-center font-black text-sm">T</div>
          <div>
            <p class="font-black text-sm">TELPP</p>
            <p class="text-xs text-gray-500">Management Profile</p>
          </div>
        </div>
      </div>
      <nav class="flex-1 p-4 space-y-1">
        <p class="text-xs text-gray-500 font-semibold tracking-widest mb-2 mt-2 px-2">OVERVIEW</p>
        <RouterLink to="/admin/dashboard"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📊 Dashboard
        </RouterLink>
        <p class="text-xs text-gray-500 font-semibold tracking-widest mb-2 mt-4 px-2">CONTENT ENGINE</p>
        <RouterLink to="/admin/news"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📰 Corporate News
        </RouterLink>
        <RouterLink to="/admin/products"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl bg-green-600 text-white text-sm font-medium">
          📦 Products
        </RouterLink>
      </nav>
      <div class="p-4 border-t border-gray-800">
        <button @click="logout" class="text-sm text-gray-400 hover:text-white transition px-2">→ Logout</button>
      </div>
    </aside>

    <!-- Main -->
    <main class="flex-1 ml-64 p-10">
      <div class="flex justify-between items-center mb-8">
        <div>
          <p class="text-xs text-gray-400 font-semibold tracking-widest">CONTENT ENGINE</p>
          <h1 class="text-3xl font-black text-gray-900">Products</h1>
        </div>
        <RouterLink to="/admin/products/create"
          class="bg-green-600 text-white px-6 py-2.5 rounded-xl text-sm font-semibold hover:bg-green-700 transition">
          + Tambah Produk
        </RouterLink>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="text-center py-20">
        <div class="w-8 h-8 border-4 border-green-600 border-t-transparent rounded-full animate-spin mx-auto"></div>
      </div>

      <!-- Empty -->
      <div v-else-if="products.length === 0"
        class="bg-white rounded-2xl p-16 text-center border border-gray-100">
        <p class="text-4xl mb-4">📦</p>
        <p class="text-gray-500 mb-4">Belum ada produk. Tambahkan produk pertama!</p>
        <RouterLink to="/admin/products/create"
          class="bg-green-600 text-white px-6 py-2 rounded-xl text-sm font-semibold hover:bg-green-700 transition">
          + Tambah Produk
        </RouterLink>
      </div>

      <!-- Product Grid -->
      <div v-else class="grid grid-cols-3 gap-6">
        <div v-for="product in products" :key="product.ID"
          class="bg-white rounded-2xl overflow-hidden border border-gray-100 shadow-sm">

          <!-- Image -->
          <div class="h-48 bg-gray-100 overflow-hidden">
            <img
              v-if="product.Images && product.Images.length > 0"
              :src="product.Images[0].image_url"
              :alt="product.name"
              class="w-full h-full object-cover"
            />
            <div v-else class="w-full h-full flex items-center justify-center">
              <span class="text-5xl">📦</span>
            </div>
          </div>

          <!-- Content -->
          <div class="p-5">
            <h3 class="font-bold text-gray-900 mb-1">{{ product.name }}</h3>
            <p class="text-gray-500 text-xs mb-4 line-clamp-2">{{ product.description }}</p>

            <!-- Image count badge -->
            <div class="flex items-center justify-between">
              <span class="bg-green-100 text-green-700 text-xs px-2 py-1 rounded-full font-semibold">
                {{ product.Images ? product.Images.length : 0 }} gambar
              </span>
              <div class="flex gap-2">
                <RouterLink :to="`/admin/products/edit/${product.ID}`"
                  class="bg-blue-50 text-blue-600 px-3 py-1.5 rounded-lg text-xs font-semibold hover:bg-blue-100 transition">
                  Edit
                </RouterLink>
                <button @click="deleteProduct(product.ID)"
                  class="bg-red-50 text-red-600 px-3 py-1.5 rounded-lg text-xs font-semibold hover:bg-red-100 transition">
                  Hapus
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import api from '../../services/api'

const router = useRouter()
const products = ref([])
const loading = ref(true)

onMounted(async () => {
  const token = localStorage.getItem('token')
  if (!token) { router.push('/admin/login'); return }
  try {
    const res = await api.get('/admin/products')
    products.value = res.data.data
  } finally {
    loading.value = false
  }
})

const deleteProduct = async (id) => {
  if (!confirm('Yakin hapus produk ini?')) return
  await api.delete(`/admin/products/${id}`)
  products.value = products.value.filter(p => p.ID !== id)
}

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/admin/login')
}
</script>