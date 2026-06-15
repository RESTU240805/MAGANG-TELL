<template>
  <div class="news-page">
    <!-- Hero Section -->
    <section class="hero">
      <div class="hero-overlay"></div>
      <div class="hero-content">
        <h1>News</h1>
        <p>
          PT Tanjung Enim Lestari Pulp and Paper is committed to world-class
          industrial standards while preserving the rich biodiversity of our
          South Sumatra home.
        </p>
      </div>
    </section>

    <!-- Main Content -->
    <div class="container">
      <!-- Loading State -->
      <div v-if="loading" class="loading-state">
        <div class="spinner"></div>
        <p>Memuat berita...</p>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="empty-state">
        <div class="empty-icon">⚠️</div>
        <p>Gagal memuat berita. Silakan coba lagi.</p>
        <button class="btn-retry" @click="fetchNews">Coba Lagi</button>
      </div>

      <!-- Empty State -->
      <div v-else-if="!newsList.length" class="empty-state">
        <div class="empty-icon">
          <svg width="60" height="60" viewBox="0 0 60 60" fill="none">
            <rect x="8" y="8" width="44" height="44" rx="4" stroke="#ccc" stroke-width="2" fill="none"/>
            <rect x="14" y="18" width="20" height="3" rx="1.5" fill="#ccc"/>
            <rect x="14" y="25" width="32" height="2" rx="1" fill="#e0e0e0"/>
            <rect x="14" y="31" width="28" height="2" rx="1" fill="#e0e0e0"/>
            <rect x="14" y="37" width="22" height="2" rx="1" fill="#e0e0e0"/>
          </svg>
        </div>
        <p>Belum ada berita.</p>
      </div>

      <!-- News Grid -->
      <div v-else>
        <!-- Featured Article (first item) -->
        <div v-if="newsList[0]" class="featured-article">
          <div class="featured-image">
            <img :src="newsList[0].image || defaultImage" :alt="newsList[0].title" />
            <span v-if="newsList[0].category" class="badge badge-community">
              {{ newsList[0].category }}
            </span>
          </div>
          <div class="featured-content">
            <div class="meta">
              <span class="meta-author">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
                  <circle cx="12" cy="7" r="4"/>
                </svg>
                Posted by {{ newsList[0].author || 'Admin' }}
              </span>
              <span class="meta-date">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
                  <line x1="16" y1="2" x2="16" y2="6"/>
                  <line x1="8" y1="2" x2="8" y2="6"/>
                  <line x1="3" y1="10" x2="21" y2="10"/>
                </svg>
                {{ formatDate(newsList[0].date) }}
              </span>
            </div>
            <h2>{{ newsList[0].title }}</h2>
            <p class="excerpt">{{ newsList[0].excerpt }}</p>
            <router-link :to="`/news/${newsList[0].id}`" class="continue-reading">
              CONTINUE READING →
            </router-link>
          </div>
        </div>

        <!-- News List (remaining items) -->
<div class="news-list">
  <div
    v-for="item in newsList.slice(1)"
    :key="item.id"
    class="news-list-item"
  >
    <div class="list-image">
      <img :src="item.image || defaultImage" :alt="item.title" />
      <span v-if="item.category" :class="['badge', getCategoryClass(item.category)]">
        {{ item.category }}
      </span>
    </div>
    <div class="list-content">
      <div class="meta">
        <span class="meta-author">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
            <circle cx="12" cy="7" r="4"/>
          </svg>
          Posted by {{ item.author || 'Admin' }}
        </span>
        <span class="meta-date">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
            <line x1="16" y1="2" x2="16" y2="6"/>
            <line x1="8" y1="2" x2="8" y2="6"/>
            <line x1="3" y1="10" x2="21" y2="10"/>
          </svg>
          {{ formatDate(item.date) }}
        </span>
      </div>
      <h2>{{ item.title }}</h2>
      <p class="excerpt">{{ item.excerpt }}</p>
      <router-link :to="`/news/${item.id}`" class="continue-reading">
        CONTINUE READING →
      </router-link>
    </div>
  </div>
</div>

        <!-- Sustainability Report Card -->
        <div class="report-section">
          <div class="report-card">
            <div class="report-card-inner">
              <div class="report-logo">
                <svg width="32" height="32" viewBox="0 0 32 32" fill="none">
                  <path d="M16 4C16 4 8 10 8 18C8 22.4 11.6 26 16 26C20.4 26 24 22.4 24 18C24 10 16 4 16 4Z" fill="white" opacity="0.8"/>
                </svg>
              </div>
              <h4>Sustainability Report 2023</h4>
            </div>
          </div>
          <div class="report-info">
            <span class="category-label sustainability">SUSTAINABILITY</span>
            <span class="card-date">12 Nov 2023</span>
            <h3>Annual Performance Review: Environmental Stewardship</h3>
            <p>Our latest audit shows a 15% reduction in water usage across the facility, exceeding our 2024 target...</p>
            <a href="#" class="view-report">VIEW REPORT ↓</a>
          </div>
        </div>

        <!-- Corporate Updates -->
        <div class="corporate-section">
          <div class="section-header">
            <h2>Corporate Updates</h2>
            <div class="nav-arrows">
              <button @click="prevPage" :disabled="currentPage <= 1" class="arrow-btn">‹</button>
              <button @click="nextPage" :disabled="currentPage >= totalPages" class="arrow-btn">›</button>
            </div>
          </div>

          <!-- Pagination -->
          <div class="pagination">
            <button @click="prevPage" :disabled="currentPage <= 1" class="page-arrow">←</button>
            <button
              v-for="page in totalPages"
              :key="page"
              @click="goToPage(page)"
              :class="['page-btn', { active: currentPage === page }]"
            >
              {{ page }}
            </button>
            <button @click="nextPage" :disabled="currentPage >= totalPages" class="page-arrow">→</button>
          </div>
        </div>

        <!-- Newsletter -->
        <section class="newsletter">
          <div class="newsletter-content">
            <h2>Stay Informed</h2>
            <p>Get the latest corporate updates and sustainability reports delivered directly to your inbox quarterly.</p>
          </div>
          <div class="newsletter-form">
            <input
              v-model="email"
              type="email"
              placeholder="Email Address"
              class="email-input"
            />
            <button @click="subscribe" class="subscribe-btn">Subscribe</button>
          </div>
        </section>
      </div>
    </div>

    <!-- Footer -->
    <footer class="footer">
      <div class="footer-container">
        <div class="footer-brand">
          <h3>TeL News</h3>
          <p>
            PT Tanjung Enim Lestari Pulp and Paper is a world-class pulp
            manufacturer operating in South Sumatra, Indonesia, with a focus on
            sustainable production and environmental stewardship.
          </p>
        </div>
        <div class="footer-col">
          <h4>OUR COMPANY</h4>
          <ul>
            <li><a href="#">Corporate Profile</a></li>
            <li><a href="#">Our Vision & Mission</a></li>
            <li><a href="#">Management Board</a></li>
            <li><a href="#">Careers at TeL</a></li>
          </ul>
        </div>
        <div class="footer-col">
          <h4>QUICK LINKS</h4>
          <ul>
            <li><a href="#">News Releases</a></li>
            <li><a href="#">Investor Relations</a></li>
            <li><a href="#">Products & Specs</a></li>
            <li><a href="#">Sustainability Center</a></li>
          </ul>
        </div>
        <div class="footer-col">
          <h4>HEAD OFFICE</h4>
          <address>
            Jl. Ir. Sutami Km. 13, Tarahan<br />
            Lampung Selatan 35352<br />
            Sumatera Selatan, Indonesia
          </address>
          <p class="phone">📞 +62 (721) 31234</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../services/api'
const newsList = ref([])
const loading = ref(true)
const error = ref(false)
const currentPage = ref(1)
const totalPages = ref(4)
const email = ref('')
const defaultImage = 'https://placehold.co/600x400/1a5c2a/white?text=TeL+News'

const formatDate = (dateStr) => {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  return date.toLocaleDateString('en-GB', { day: '2-digit', month: 'short', year: 'numeric' })
}

const getCategoryClass = (category) => {
  if (!category) return 'default'
  const cat = category.toLowerCase()
  if (cat.includes('community')) return 'community'
  if (cat.includes('sport') || cat.includes('csr')) return 'sport-csr'
  if (cat.includes('corporate')) return 'corporate'
  if (cat.includes('sustain')) return 'sustainability'
  return 'default'
}

const fetchNews = async () => {
  loading.value = true
  error.value = false
  try {
    const res = await api.get(`/news?page=${currentPage.value}`)
    const data = res.data

    const rawList = data.data || data.news || data || []
    newsList.value = rawList
      .filter(item => item.is_published)
      .map(item => ({
        id: item.ID || item.id,
        title: item.title,
        excerpt: item.summary || '',
        category: item.category,
        author: item.author || 'Admin',
        date: item.CreatedAt || item.created_at || item.date,
        image: item.Images?.[0]?.image_url || '',
      }))

    if (data.total_pages) totalPages.value = data.total_pages

  } catch (e) {
    console.error(e)
    error.value = true
  } finally {
    loading.value = false
  }
}



const goToPage = (page) => {
  currentPage.value = page
  fetchNews()
}

const prevPage = () => {
  if (currentPage.value > 1) goToPage(currentPage.value - 1)
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) goToPage(currentPage.value + 1)
}

const subscribe = () => {
  if (!email.value) return
  alert(`Terima kasih! Email ${email.value} telah didaftarkan.`)
  email.value = ''
}

onMounted(() => {
  fetchNews()
})
</script>

<style scoped>
/* ─── Reset & Base ─────────────────────────────────────────── */
* { box-sizing: border-box; }

.news-page {
  font-family: 'Segoe UI', Arial, sans-serif;
  color: #1a1a1a;
  background: #fff;
}

/* ─── Hero ─────────────────────────────────────────────────── */
.hero {
  position: relative;
  height: 320px;
  background: url('https://placehold.co/1400x400/1a3a2a/white?text=') center/cover no-repeat;
  background-color: #2d5a3d;
  display: flex;
  align-items: flex-end;
}

.hero-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(to right, rgba(0,0,0,0.55) 0%, rgba(0,0,0,0.1) 60%);
}

.hero-content {
  position: relative;
  z-index: 1;
  padding: 40px 48px;
  max-width: 520px;
}

.hero-content h1 {
  font-size: 2.8rem;
  font-weight: 700;
  color: #fff;
  margin: 0 0 12px;
}

.hero-content p {
  font-size: 0.95rem;
  color: rgba(255,255,255,0.88);
  line-height: 1.6;
  margin: 0;
}

/* ─── Container ────────────────────────────────────────────── */
.container {
  max-width: 1100px;
  margin: 0 auto;
  padding: 48px 24px;
}

/* ─── Loading & Empty ──────────────────────────────────────── */
.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 80px 0;
  gap: 16px;
  color: #888;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 3px solid #e0e0e0;
  border-top-color: #2d7a3d;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

@keyframes spin { to { transform: rotate(360deg); } }

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 80px 0;
  gap: 12px;
  color: #aaa;
}

.empty-icon { margin-bottom: 8px; }

.btn-retry {
  margin-top: 8px;
  padding: 8px 20px;
  background: #2d7a3d;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

/* ─── Featured Article ─────────────────────────────────────── */
.featured-article {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 0;
  border: 1px solid #e8e8e8;
  border-radius: 4px;
  overflow: hidden;
  margin-bottom: 32px;
}

.featured-image {
  position: relative;
  min-height: 280px;
  background: #2d5a3d;
}

.featured-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.featured-content {
  padding: 32px 28px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.meta {
  display: flex;
  gap: 16px;
  font-size: 0.78rem;
  color: #777;
  margin-bottom: 12px;
}

.meta-author, .meta-date {
  display: flex;
  align-items: center;
  gap: 4px;
}

.featured-content h2 {
  font-size: 1.4rem;
  font-weight: 700;
  color: #1a1a1a;
  margin: 0 0 12px;
  line-height: 1.4;
}

.excerpt {
  font-size: 0.9rem;
  color: #555;
  line-height: 1.6;
  margin: 0 0 20px;
}

.continue-reading {
  color: #2d7a3d;
  font-size: 0.8rem;
  font-weight: 700;
  letter-spacing: 0.05em;
  text-decoration: none;
}

.continue-reading:hover { text-decoration: underline; }

/* ─── Badges ───────────────────────────────────────────────── */
.badge {
  position: absolute;
  top: 12px;
  left: 12px;
  padding: 4px 10px;
  font-size: 0.68rem;
  font-weight: 700;
  letter-spacing: 0.06em;
  text-transform: uppercase;
  border-radius: 3px;
  color: #fff;
}

.badge.community, .badge-community { background: #2d7a3d; }
.badge.sport-csr { background: #e67e22; }
.badge.corporate { background: #2c3e50; }
.badge.sustainability { background: #16a085; }
.badge.default { background: #666; }

/* ─── News Grid ────────────────────────────────────────────── */
.news-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 24px;
  margin-bottom: 40px;
}

.news-card {
  border: 1px solid #e8e8e8;
  border-radius: 4px;
  overflow: hidden;
  transition: box-shadow 0.2s;
}

.news-card:hover { box-shadow: 0 4px 20px rgba(0,0,0,0.1); }

.card-image {
  position: relative;
  height: 180px;
  background: #2d5a3d;
}

.card-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.card-body { padding: 18px; }

.card-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.category-label {
  font-size: 0.68rem;
  font-weight: 700;
  letter-spacing: 0.06em;
  text-transform: uppercase;
}

.category-label.community { color: #2d7a3d; }
.category-label.sport-csr { color: #e67e22; }
.category-label.corporate { color: #2c3e50; }
.category-label.sustainability { color: #16a085; }
.category-label.default { color: #666; }

.card-date { font-size: 0.72rem; color: #999; }

.news-card h3 {
  font-size: 1rem;
  font-weight: 700;
  color: #1a1a1a;
  margin: 0 0 8px;
  line-height: 1.4;
}

.card-excerpt {
  font-size: 0.85rem;
  color: #666;
  line-height: 1.55;
  margin: 0 0 14px;
}

.read-article {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  font-size: 0.75rem;
  font-weight: 700;
  color: #2d7a3d;
  text-decoration: none;
  letter-spacing: 0.04em;
}

.read-article:hover { text-decoration: underline; }

/* ─── Report Section ───────────────────────────────────────── */
.report-section {
  display: grid;
  grid-template-columns: 240px 1fr;
  gap: 24px;
  align-items: start;
  margin-bottom: 48px;
  padding: 24px;
  border: 1px solid #e8e8e8;
  border-radius: 4px;
}

.report-card {
  background: linear-gradient(135deg, #2d7a3d 0%, #1a5c2a 100%);
  border-radius: 6px;
  padding: 28px 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 130px;
}

.report-card-inner {
  text-align: center;
  color: #fff;
}

.report-card-inner h4 {
  font-size: 1rem;
  font-weight: 700;
  margin: 12px 0 0;
}

.report-info h3 {
  font-size: 1.1rem;
  font-weight: 700;
  margin: 10px 0 8px;
  line-height: 1.4;
}

.report-info p { font-size: 0.88rem; color: #666; margin: 0 0 12px; }

.view-report {
  font-size: 0.78rem;
  font-weight: 700;
  color: #2d7a3d;
  text-decoration: none;
}

/* ─── Corporate / Pagination ───────────────────────────────── */
.corporate-section { margin-bottom: 48px; }

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  padding-bottom: 12px;
  border-bottom: 3px solid #2d7a3d;
}

.section-header h2 {
  font-size: 1.5rem;
  font-weight: 700;
  margin: 0;
}

.nav-arrows { display: flex; gap: 8px; }

.arrow-btn {
  width: 34px;
  height: 34px;
  border: 1px solid #ddd;
  background: #fff;
  border-radius: 50%;
  cursor: pointer;
  font-size: 1.1rem;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}

.arrow-btn:hover:not(:disabled) { background: #2d7a3d; color: #fff; border-color: #2d7a3d; }
.arrow-btn:disabled { opacity: 0.35; cursor: not-allowed; }

.pagination {
  display: flex;
  align-items: center;
  gap: 6px;
}

.page-btn {
  width: 36px;
  height: 36px;
  border: 1px solid #ddd;
  background: #fff;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9rem;
  transition: all 0.2s;
}

.page-btn.active {
  background: #2d7a3d;
  color: #fff;
  border-color: #2d7a3d;
}

.page-btn:hover:not(.active) { background: #f0f8f2; border-color: #2d7a3d; }

.page-arrow {
  padding: 6px 10px;
  border: 1px solid #ddd;
  background: #fff;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9rem;
}

.page-arrow:disabled { opacity: 0.35; cursor: not-allowed; }

/* ─── Newsletter ───────────────────────────────────────────── */
.newsletter {
  background: #f5f5f5;
  border-radius: 6px;
  padding: 40px 36px;
  display: grid;
  grid-template-columns: 1fr auto;
  gap: 32px;
  align-items: center;
  margin-bottom: 0;
}

.newsletter-content h2 {
  font-size: 1.5rem;
  font-weight: 700;
  margin: 0 0 8px;
}

.newsletter-content p {
  font-size: 0.9rem;
  color: #555;
  margin: 0;
  max-width: 460px;
}

.newsletter-form {
  display: flex;
  gap: 0;
  min-width: 320px;
}

.email-input {
  flex: 1;
  padding: 12px 16px;
  border: 1px solid #ddd;
  border-right: none;
  border-radius: 4px 0 0 4px;
  font-size: 0.9rem;
  outline: none;
}

.email-input:focus { border-color: #2d7a3d; }

.subscribe-btn {
  padding: 12px 24px;
  background: #2d7a3d;
  color: #fff;
  border: none;
  border-radius: 0 4px 4px 0;
  font-size: 0.9rem;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s;
}

.subscribe-btn:hover { background: #1a5c2a; }

/* ─── Footer ───────────────────────────────────────────────── */
.footer {
  background: #1a1a1a;
  color: #ccc;
  padding: 48px 0 24px;
  margin-top: 0;
}

.footer-container {
  max-width: 1100px;
  margin: 0 auto;
  padding: 0 24px;
  display: grid;
  grid-template-columns: 2fr 1fr 1fr 1.5fr;
  gap: 40px;
}

.footer-brand h3 {
  color: #fff;
  font-size: 1.1rem;
  font-weight: 700;
  margin: 0 0 12px;
  border-bottom: 2px solid #2d7a3d;
  padding-bottom: 8px;
  display: inline-block;
}

.footer-brand p { font-size: 0.85rem; line-height: 1.65; color: #aaa; }

.footer-col h4 {
  color: #fff;
  font-size: 0.78rem;
  font-weight: 700;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  margin: 0 0 16px;
  padding-bottom: 8px;
  border-bottom: 1px solid #333;
}

.footer-col ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.footer-col ul li { margin-bottom: 8px; }

.footer-col ul li a {
  color: #aaa;
  text-decoration: none;
  font-size: 0.85rem;
  transition: color 0.2s;
}

.footer-col ul li a:hover { color: #2d7a3d; }

.footer-col address {
  font-style: normal;
  font-size: 0.85rem;
  line-height: 1.7;
  color: #aaa;
}

.phone { font-size: 0.85rem; color: #aaa; margin-top: 8px; }

/* ─── Responsive ───────────────────────────────────────────── */
@media (max-width: 768px) {
  .featured-article {
    grid-template-columns: 1fr;
  }

  .featured-image { min-height: 200px; }

  .news-grid { grid-template-columns: 1fr; }

  .report-section {
    grid-template-columns: 1fr;
  }

  .newsletter {
    grid-template-columns: 1fr;
    gap: 20px;
  }

  .newsletter-form { min-width: auto; }

  .footer-container {
    grid-template-columns: 1fr 1fr;
    gap: 28px;
  }

  .hero-content { padding: 24px; }
  .hero-content h1 { font-size: 2rem; }
}

@media (max-width: 480px) {
  .footer-container { grid-template-columns: 1fr; }
}

/* ─── News List ────────────────────────────────────────────── */
.news-list {
  display: flex;
  flex-direction: column;
  margin-bottom: 40px;
}

.news-list-item {
  display: grid;
  grid-template-columns: 320px 1fr;
  gap: 28px;
  padding: 28px 0;
  border-bottom: 1px solid #e8e8e8;
  align-items: flex-start;
}

.list-image {
  position: relative;
  height: 220px;
  background: #2d5a3d;
  flex-shrink: 0;
}

.list-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.list-content {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  padding-top: 4px;
}

.list-content h2 {
  font-size: 1.3rem;
  font-weight: 700;
  color: #1a1a1a;
  margin: 0 0 12px;
  line-height: 1.4;
}

.list-content .excerpt {
  font-size: 0.9rem;
  color: #555;
  line-height: 1.7;
  margin: 0 0 16px;
}

.list-content .continue-reading {
  display: inline-block;
  border: 1px solid #ccc;
  padding: 8px 20px;
  font-size: 0.75rem;
  font-weight: 600;
  letter-spacing: 0.08em;
  color: #444;
  text-decoration: none;
  width: fit-content;
}

.list-content .continue-reading:hover {
  background: #f5f5f5;
  text-decoration: none;
}

@media (max-width: 768px) {
  .news-list-item {
    grid-template-columns: 1fr;
  }
  .list-image { height: 200px; }
}
</style>