<template>
  <div class="min-h-screen bg-gray-100 flex">

    <!-- ─── Sidebar ─────────────────────────────────────────── -->
    <aside class="w-64 bg-gray-950 text-white flex flex-col fixed h-full z-40">
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
        <RouterLink to="/admin/slider"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🖼️ Product Slider
        </RouterLink>
        <RouterLink to="/admin/product-page"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📝 Product Page
        </RouterLink>
        <RouterLink to="/admin/about"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🏢 About Section
        </RouterLink>
        <RouterLink to="/admin/team-members"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          👥 Our Team
        </RouterLink>
      </nav>
      <div class="p-4 border-t border-gray-800">
        <button @click="logout" class="text-sm text-gray-400 hover:text-white transition px-2">→ Logout</button>
      </div>
    </aside>

    <!-- ─── Main ─────────────────────────────────────────────── -->
    <main class="flex-1 ml-64 p-10">

      <!-- Header -->
      <div class="flex justify-between items-center mb-8">
        <div>
          <p class="text-xs text-gray-400 font-semibold tracking-widest">CONTENT ENGINE</p>
          <h1 class="text-3xl font-black text-gray-900">Products</h1>
        </div>
        <button @click="openCreate"
          class="bg-green-600 text-white px-6 py-2.5 rounded-xl text-sm font-semibold hover:bg-green-700 transition flex items-center gap-2">
          <span class="text-lg leading-none">+</span> Tambah Produk
        </button>
      </div>

      <!-- Alert -->
      <div v-if="alert.show"
        :class="['mb-4 px-4 py-3 rounded-xl text-sm font-medium flex items-center gap-2',
          alert.type === 'success' ? 'bg-green-50 text-green-700 border border-green-200' : 'bg-red-50 text-red-700 border border-red-200']">
        {{ alert.type === 'success' ? '✅' : '❌' }} {{ alert.message }}
      </div>

      <!-- Loading -->
      <div v-if="loading" class="flex items-center justify-center py-20 gap-3 text-gray-400">
        <div class="w-5 h-5 border-2 border-gray-200 border-t-green-500 rounded-full animate-spin"></div>
        Memuat data...
      </div>

      <!-- Empty -->
      <div v-else-if="products.length === 0"
        class="bg-white rounded-2xl p-16 text-center border border-gray-100">
        <p class="text-4xl mb-4">📦</p>
        <p class="text-gray-500 mb-4">Belum ada produk.</p>
        <button @click="openCreate"
          class="bg-green-600 text-white px-6 py-2 rounded-xl text-sm font-semibold hover:bg-green-700 transition">
          + Tambah Produk
        </button>
      </div>

      <!-- Table -->
      <div v-else class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-100">
            <tr>
              <th class="text-left px-6 py-4 text-xs font-semibold text-gray-500 tracking-widest">PRODUK</th>
              <th class="text-left px-6 py-4 text-xs font-semibold text-gray-500 tracking-widest">KATEGORI</th>
              <th class="text-left px-6 py-4 text-xs font-semibold text-gray-500 tracking-widest">STATUS</th>
              <th class="text-left px-6 py-4 text-xs font-semibold text-gray-500 tracking-widest">AKSI</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-50">
            <tr v-for="p in products" :key="p.ID" class="hover:bg-gray-50 transition">
              <td class="px-6 py-4">
                <div class="flex items-center gap-3">
                  <img v-if="p.thumbnail_path"
                    :src="getImageUrl(p.thumbnail_path)"
                    class="w-12 h-12 rounded-lg object-cover flex-shrink-0 bg-gray-100"
                    @error="(e) => e.target.style.display='none'" />
                  <div class="w-12 h-12 rounded-lg bg-gray-100 flex items-center justify-center flex-shrink-0 text-xl" v-else>📦</div>
                  <div>
                    <p class="font-semibold text-gray-900 text-sm">{{ p.name }}</p>
                    <p class="text-gray-400 text-xs mt-0.5 line-clamp-1">{{ p.summary }}</p>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <span class="bg-green-100 text-green-700 px-3 py-1 rounded-full text-xs font-semibold">
                  {{ p.category || 'General' }}
                </span>
              </td>
              <td class="px-6 py-4">
                <span :class="['px-3 py-1 rounded-full text-xs font-semibold',
                  p.is_active ? 'bg-green-100 text-green-700' : 'bg-gray-100 text-gray-500']">
                  {{ p.is_active ? 'Active' : 'Draft' }}
                </span>
              </td>
              <td class="px-6 py-4">
                <div class="flex gap-2">
                  <button @click="openEdit(p)"
                    class="bg-blue-50 text-blue-600 px-3 py-1.5 rounded-lg text-xs font-semibold hover:bg-blue-100 transition">
                    Edit
                  </button>
                  <button @click="confirmDelete(p)"
                    class="bg-red-50 text-red-600 px-3 py-1.5 rounded-lg text-xs font-semibold hover:bg-red-100 transition">
                    Hapus
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </main>

    <!-- ─── Modal Form ───────────────────────────────────────── -->
    <Transition name="modal">
      <div v-if="showModal"
        class="fixed inset-0 z-50 flex items-center justify-center p-4"
        @click.self="closeModal">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-2xl max-h-[90vh] overflow-y-auto">

          <div class="flex items-center justify-between p-6 border-b border-gray-100">
            <h2 class="text-xl font-black text-gray-900">
              {{ isEdit ? '✏️ Edit Produk' : '➕ Tambah Produk' }}
            </h2>
            <button @click="closeModal" class="text-gray-400 hover:text-gray-600 text-2xl leading-none">×</button>
          </div>

          <form @submit.prevent="submitForm" class="p-6 space-y-5">

            <!-- Nama -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Nama Produk <span class="text-red-500">*</span></label>
              <input v-model="form.name" type="text" placeholder="Masukkan nama produk..." required
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400" />
            </div>

            <!-- Kategori -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Kategori</label>
              <div class="flex gap-2 flex-wrap mb-2">
                <button v-for="cat in categories" :key="cat" type="button"
                  @click="form.category = cat"
                  :class="['px-3 py-1 rounded-full text-xs font-semibold border transition',
                    form.category === cat ? 'bg-green-600 text-white border-green-600' : 'bg-white text-gray-500 border-gray-200 hover:border-green-400']">
                  {{ cat }}
                </button>
              </div>
              <input v-model="form.category" type="text" placeholder="Atau ketik kategori sendiri..."
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400" />
            </div>

            <!-- Summary -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Ringkasan</label>
              <textarea v-model="form.summary" rows="2" placeholder="Deskripsi singkat produk (tampil di card)..."
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400 resize-none"></textarea>
            </div>

            <!-- Description -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Deskripsi Lengkap <span class="text-red-500">*</span></label>
              <textarea v-model="form.description" rows="4" placeholder="Deskripsi lengkap produk..." required
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400 resize-none"></textarea>
            </div>

            <!-- Tags -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">
                Tags / Spesifikasi
                <span class="text-xs text-gray-400 font-normal ml-1">(pisahkan dengan koma)</span>
              </label>
              <input v-model="form.tags" type="text" placeholder="High Strength, Bleached, Premium..."
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400" />
              <!-- Preview tags -->
              <div class="flex flex-wrap gap-1.5 mt-2" v-if="form.tags">
                <span v-for="tag in form.tags.split(',').filter(t=>t.trim())" :key="tag"
                  class="text-xs bg-green-100 text-green-700 px-2.5 py-1 rounded-full font-semibold">
                  {{ tag.trim() }}
                </span>
              </div>
            </div>

            <!-- Thumbnail Upload -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Thumbnail</label>
              <div class="border-2 border-dashed border-gray-200 rounded-xl p-4 text-center hover:border-green-400 transition cursor-pointer"
                @click="$refs.thumbInput.click()"
                @dragover.prevent @drop.prevent="handleDrop">
                <input ref="thumbInput" type="file" accept="image/*" class="hidden" @change="handleThumbChange" />
                <div v-if="thumbPreview || form.thumbnail_path">
                  <img :src="thumbPreview || getImageUrl(form.thumbnail_path)"
                    class="mx-auto h-32 object-cover rounded-lg mb-2" />
                  <p class="text-xs text-gray-400">Klik untuk ganti gambar</p>
                </div>
                <div v-else>
                  <p class="text-3xl mb-2">🖼️</p>
                  <p class="text-sm text-gray-500">Klik atau drag & drop gambar</p>
                  <p class="text-xs text-gray-400 mt-1">PNG, JPG, JPEG (maks 5MB)</p>
                </div>
              </div>
            </div>

            <!-- Active toggle -->
            <div class="flex items-center justify-between bg-gray-50 rounded-xl px-4 py-3">
              <div>
                <p class="text-sm font-semibold text-gray-700">Aktifkan Produk</p>
                <p class="text-xs text-gray-400 mt-0.5">Produk akan tampil di halaman publik</p>
              </div>
              <button type="button" @click="form.is_active = !form.is_active"
                :class="['relative w-12 h-6 rounded-full transition-colors duration-200',
                  form.is_active ? 'bg-green-500' : 'bg-gray-300']">
                <span :class="['absolute top-0.5 left-0.5 w-5 h-5 bg-white rounded-full shadow transition-transform duration-200',
                  form.is_active ? 'translate-x-6' : 'translate-x-0']"></span>
              </button>
            </div>

            <!-- Error -->
            <div v-if="formError" class="bg-red-50 text-red-600 text-sm px-4 py-3 rounded-xl border border-red-200">
              ❌ {{ formError }}
            </div>

            <!-- Actions -->
            <div class="flex gap-3 pt-2">
              <button type="button" @click="closeModal"
                class="flex-1 border border-gray-200 text-gray-600 py-2.5 rounded-xl text-sm font-semibold hover:bg-gray-50 transition">
                Batal
              </button>
              <button type="submit" :disabled="submitting"
                class="flex-1 bg-green-600 text-white py-2.5 rounded-xl text-sm font-semibold hover:bg-green-700 transition disabled:opacity-60 flex items-center justify-center gap-2">
                <div v-if="submitting" class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
                {{ submitting ? 'Menyimpan...' : (isEdit ? 'Simpan Perubahan' : 'Tambah Produk') }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </Transition>

    <!-- ─── Modal Hapus ──────────────────────────────────────── -->
    <Transition name="modal">
      <div v-if="showDeleteModal"
        class="fixed inset-0 z-50 flex items-center justify-center p-4"
        @click.self="showDeleteModal = false">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-sm p-6 text-center">
          <p class="text-4xl mb-3">🗑️</p>
          <h3 class="text-lg font-black text-gray-900 mb-1">Hapus Produk?</h3>
          <p class="text-sm text-gray-500 mb-6">
            "<span class="font-semibold text-gray-700">{{ selectedProduct?.name }}</span>"
            akan dihapus permanen.
          </p>
          <div class="flex gap-3">
            <button @click="showDeleteModal = false"
              class="flex-1 border border-gray-200 text-gray-600 py-2.5 rounded-xl text-sm font-semibold hover:bg-gray-50 transition">
              Batal
            </button>
            <button @click="doDelete" :disabled="submitting"
              class="flex-1 bg-red-600 text-white py-2.5 rounded-xl text-sm font-semibold hover:bg-red-700 transition disabled:opacity-60">
              {{ submitting ? 'Menghapus...' : 'Hapus' }}
            </button>
          </div>
        </div>
      </div>
    </Transition>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import api from '../../services/api'

const BASE_URL = 'http://localhost:8080'
const router = useRouter()

const products       = ref([])
const loading        = ref(true)
const submitting     = ref(false)
const showModal      = ref(false)
const showDeleteModal= ref(false)
const isEdit         = ref(false)
const selectedProduct= ref(null)
const thumbPreview   = ref(null)
const thumbFile      = ref(null)
const formError      = ref('')
const alert          = ref({ show: false, type: 'success', message: '' })

const categories = ['Pulp', 'Tissue', 'Paper', 'Specialty', 'Chemical']

const form = ref({
  name: '', category: '', summary: '', description: '',
  tags: '', thumbnail_path: '', is_active: true,
})

const getImageUrl = (path) => {
  if (!path) {return ''}
  if (path.startsWith('http')) {return path}
  return `${BASE_URL}/${path.replace(/^\//, '')}`
}

const showAlert = (type, message) => {
  alert.value = { show: true, type, message }
  setTimeout(() => { alert.value.show = false }, 3000)
}

const fetchProducts = async () => {
  loading.value = true
  try {
    const res = await api.get('/admin/products')
    products.value = res.data?.data || res.data || []
  } catch { showAlert('error', 'Gagal memuat produk') }
  finally { loading.value = false }
}

const openCreate = () => {
  isEdit.value = false
  form.value = { name: '', category: '', summary: '', description: '', tags: '', thumbnail_path: '', is_active: true }
  thumbPreview.value = null; thumbFile.value = null; formError.value = ''
  showModal.value = true
}

const openEdit = (p) => {
  isEdit.value = true; selectedProduct.value = p
  const tagsStr = Array.isArray(p.tags) ? p.tags.join(', ') : (p.tags || '')
  form.value = {
    name: p.name || '', category: p.category || '',
    summary: p.summary || '', description: p.description || '',
    tags: tagsStr, thumbnail_path: p.thumbnail_path || '',
    is_active: p.is_active !== false,
  }
  thumbPreview.value = null; thumbFile.value = null; formError.value = ''
  showModal.value = true
}

const closeModal = () => { showModal.value = false; formError.value = '' }

const handleThumbChange = (e) => {
  const file = e.target.files[0]; if (!file) {return}
  if (file.size > 5 * 1024 * 1024) { formError.value = 'Gambar maks 5MB'; return }
  thumbFile.value = file
  thumbPreview.value = URL.createObjectURL(file)
}

const handleDrop = (e) => {
  const file = e.dataTransfer.files[0]
  if (file) { thumbFile.value = file; thumbPreview.value = URL.createObjectURL(file) }
}

const uploadThumb = async () => {
  if (!thumbFile.value) {return form.value.thumbnail_path}
  const fd = new FormData()
  fd.append('image', thumbFile.value)
  const res = await api.post('/upload', fd, { headers: { 'Content-Type': 'multipart/form-data' } })
  return res.data?.url || ''
}

const submitForm = async () => {
  if (!form.value.name.trim() || !form.value.description.trim()) {
    formError.value = 'Nama dan deskripsi wajib diisi'; return
  }
  submitting.value = true; formError.value = ''
  try {
    if (thumbFile.value) {form.value.thumbnail_path = await uploadThumb()}
    const tagsArr = form.value.tags
      ? form.value.tags.split(',').map(t => t.trim()).filter(Boolean)
      : []
    const payload = {
      name: form.value.name,
      category: form.value.category,
      summary: form.value.summary,
      description: form.value.description,
      tags: tagsArr.join(', '),
      thumbnail_path: form.value.thumbnail_path,
      is_active: form.value.is_active,
    }
    if (isEdit.value) {
      await api.put(`/admin/products/${selectedProduct.value.ID}`, payload)
      showAlert('success', 'Produk berhasil diperbarui!')
    } else {
      await api.post('/admin/products', payload)
      showAlert('success', 'Produk berhasil ditambahkan!')
    }
    closeModal(); await fetchProducts()
  } catch (e) {
    formError.value = e.response?.data?.message || 'Terjadi kesalahan. Coba lagi.'
  } finally { submitting.value = false }
}

const confirmDelete = (p) => { selectedProduct.value = p; showDeleteModal.value = true }
const doDelete = async () => {
  submitting.value = true
  try {
    await api.delete(`/admin/products/${selectedProduct.value.ID}`)
    products.value = products.value.filter(p => p.ID !== selectedProduct.value.ID)
    showDeleteModal.value = false
    showAlert('success', 'Produk berhasil dihapus!')
  } catch { showAlert('error', 'Gagal menghapus produk') }
  finally { submitting.value = false }
}

const logout = () => {
  localStorage.removeItem('token'); localStorage.removeItem('user')
  router.push('/admin/login')
}

onMounted(() => {
  if (!localStorage.getItem('token')) { router.push('/admin/login'); return }
  fetchProducts()
})
</script>

<style scoped>
.modal-enter-active, .modal-leave-active { transition: opacity 0.2s ease; }
.modal-enter-from, .modal-leave-to       { opacity: 0; }
</style>