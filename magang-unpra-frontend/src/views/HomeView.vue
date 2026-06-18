<template>
  <div class="page">

    <section class="hero">
      <img
        v-for="(item, index) in heroSlides"
        :key="item.src"
        :src="item.src"
        alt="Pulp and Paper Factory"
        class="hero-bg"
        :class="{ 'hero-bg--active': index === currentSlide }"
      />
      <div class="hero-overlay"></div>
      <div class="hero-content">
        <transition name="fade-title" mode="out-in">
          <div
            :key="currentSlide"
            class="hero-title-block"
          >
            <span
              v-if="heroSlides[currentSlide].sub"
              class="hero-sub"
            >{{ heroSlides[currentSlide].sub }}</span>
            <h1 :class="{ 'hero-quote': heroSlides[currentSlide].type === 'quote' }">
              <span class="char-container">
                <span
                  v-for="(c, i) in currentChars"
                  :key="i"
                  class="char"
                  :class="{
                    'accent': c.type === 'accent',
                    'char-space': c.text === ' '
                  }"
                  :style="{ animationDelay: c.delay + 'ms' }"
                >{{ c.text }}</span>
              </span>
            </h1>
          </div>
        </transition>
      </div>
      <button class="scroll-down" aria-label="Scroll down">
        <span class="arrow-down">↓</span>
      </button>
    </section>

    <section class="about reveal reveal-left">
      <div class="about-inner">
        <div class="about-text">
          <span class="eyebrow">About TELPP</span>
          <p>
            PT Tanjungenim Lestari Pulp and Paper ( TELPP) is world class manufacturer of high product quality and environmental friendly market pulp mill. This was established on June 18, 1990, commenced construction in mid-1997 and the commercial operation started on May, 2000 . The mill is located in 1,250 ha area in the Banuayu village, District Empat Petulai Dangku, Muara Enim Regency, South Sumatra province, Indonesia.
          </p>
          
        </div>

        <div class="about-media reveal reveal-right" data-delay="200">
          <img src="/images/gedung.jpeg" alt="TELPP Worker" class="about-img" />
          <div class="about-badge">
            <span class="badge-number">20+</span>
            <span class="badge-label">Years of Excellence</span>
          </div>
        </div>
      </div>
    </section>

    <section class="community reveal reveal-up">
      <div class="section-heading">
        <h2>Local Community Development</h2>
        <span class="heading-underline"></span>
      </div>

      <div class="community-grid">
        <router-link to="/sustainability/csr/community" class="community-card reveal reveal-up" data-delay="0">
          <div class="icon-box icon-pink">
            <img src="/images/gedung.jpeg" alt="Education icon" />
          </div>
          <h3>Education</h3>
          <p>
            Prepare a generation of quality sources from internal employees as well as from the village community around the company of Telpp. We achieved this through scholarships, traineeships, Development of Community Education, School foundation program and school facilities & infrastructure improvement.
          </p>
        </router-link>

        <router-link to="/sustainability/csr/community" class="community-card reveal reveal-up" data-delay="150">
          <div class="icon-box icon-green">
            <img src="/images/gedung.jpeg" alt="Infrastructure icon" />
          </div>
          <h3>Infrastructure</h3>
          <p>
            Actively participating to assist availability of public
            facilities as the driving force of economic growth and
            improving the quality of life and welfare of local communities
            in the village around the company PT Tanjungenim Lestari Pulp
            and Paper.
          </p>
        </router-link>

        <router-link to="/sustainability/csr/community" class="community-card reveal reveal-up" data-delay="300">
          <div class="icon-box icon-teal">
            <img src="/images/gedung.jpeg" alt="Local economy icon" />
          </div>
          <h3>Local Economy Development</h3>
          <p>
            Our beneficiary community businesses of Small, Medium Enterprise
            Development that have an impact on the direct, indirect and the
            ability of people involved in order to become economically
            independent and sustainable primarily around the company of PT
            Tanjungenim Lestari Pulp and Paper.
          </p>
        </router-link>

        <router-link to="/sustainability/csr/community" class="community-card reveal reveal-up" data-delay="450">
          <div class="icon-box icon-blue">
            <img src="/images/gedung.jpeg" alt="Health icon" />
          </div>
          <h3>Health &amp; Conservation</h3>
          <p>
            Increasing community awareness, changing the mindset of
            not stakeholders and the community in health development by
            increasing community empowerment efforts through quality
            appropriate service facilities and community movements in
            healthy living.
          </p>
        </router-link>
      </div>
    </section>

    <section class="sustainability reveal reveal-up">
      <div class="section-heading section-heading--light">
        <h2>Sustainability</h2>
        <span class="heading-underline heading-underline--light"></span>
      </div>

      <div class="sustainability-grid">
        <div 
          v-for="(item, index) in sustainabilityItems" 
          :key="index" 
          class="sustainability-item reveal reveal-up" 
          :data-delay="index * 100"
        >
          <img :src="item.img" :alt="item.title" />
          <div class="sustainability-overlay">
            <div class="overlay-content">
              <h3>{{ item.title }}</h3>
              <span class="overlay-divider"></span>
              <router-link :to="item.link" class="search-btn" :aria-label="'Search ' + item.title">
                <span class="search-icon">🔍</span>
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section class="product reveal reveal-left">
      <div class="product-inner">
        <div class="product-text">
          <span class="eyebrow">Why Choose Pellita</span>
          <h2>
            100% Pellita Pulp.<br />
            The only one and the 1st<br />
            market pulp in the world
          </h2>
          <p class="product-sub">with Excellent Smoothness, Softness and Strength</p>

          <div class="accordion">
            <div class="accordion-group">
              <button 
                class="accordion-item" 
                :class="{ 'accordion-item--active': activeAccordion === 'certified' }"
                @click="toggleAccordion('certified')"
              >
                Certified
                <span class="accordion-icon">{{ activeAccordion === 'certified' ? '⌃' : '⌄' }}</span>
              </button>
              <div v-if="activeAccordion === 'certified'" class="accordion-panel">
                <img src="/images/hutan3.jpeg" alt="Certified Content" class="accordion-panel-img" />
              </div>
            </div>

            <div class="accordion-group">
              <button 
                class="accordion-item" 
                :class="{ 'accordion-item--active': activeAccordion === 'compliance' }"
                @click="toggleAccordion('compliance')"
              >
                Compliance Management
                <span class="accordion-icon">{{ activeAccordion === 'compliance' ? '⌃' : '⌄' }}</span>
              </button>
              <div v-if="activeAccordion === 'compliance'" class="accordion-panel">
                <img src="/images/people3.jpeg" alt="Compliance Content" class="accordion-panel-img" />
              </div>
            </div>

            <div class="accordion-group">
              <button 
                class="accordion-item" 
                :class="{ 'accordion-item--active': activeAccordion === 'concept' }"
                @click="toggleAccordion('concept')"
              >
                The Concept
                <span class="accordion-icon">{{ activeAccordion === 'concept' ? '⌃' : '⌄' }}</span>
              </button>
              <div v-if="activeAccordion === 'concept'" class="accordion-panel">
                <img src="/images/pulp.jpeg" alt="The Concept Content" class="accordion-panel-img" />
              </div>
            </div>
          </div>
        </div>

        <div class="product-media reveal reveal-right" data-delay="200">
          <img src="/images/kimia.jpeg" alt="Pellita pulp sheet texture" />
        </div>
      </div>
    </section>

    <section class="news reveal reveal-up">
      <div class="news-header">
        <h2>Latest news</h2>
        <RouterLink to="/news" class="view-all">More All News <span class="arrow">→</span></RouterLink>
      </div>

      <div class="news-grid">
        <article v-for="(item, i) in latestNews" :key="item.id"
          class="news-card reveal reveal-up" :data-delay="i * 150">
          <RouterLink :to="`/news/${item.id}`" class="block">
            <img :src="getImageUrl(item.image)" :alt="item.title" class="news-thumb" />
          </RouterLink>
          <div class="news-body">
            <span class="news-meta">
              <span class="news-author">{{ item.author }}</span>
              <span class="news-date">{{ formatDate(item.date) }}</span>
            </span>
            <RouterLink :to="`/news/${item.id}`">
              <h3 class="hover:text-green-600 transition-colors cursor-pointer">{{ item.title }}</h3>
            </RouterLink>
            <p>{{ item.excerpt }}</p>
          </div>
        </article>

        <div v-if="latestNews.length === 0" class="col-span-2 text-center text-gray-400 py-10">
          Loading news...
        </div>
      </div>
    </section>

   <section class="offices reveal reveal-up">
      <div class="section-heading">
        <h2>Our Offices</h2>
        <span class="heading-underline"></span>
      </div>

      <div class="offices-grid">
        <div class="office-item reveal reveal-up" data-delay="0">
          <img src="/images/jakarta.jpeg" alt="Jakarta Office" />
          <h3 class="office-title">Jakarta</h3>
          <p class="office-address">Menara Astra 22nd floor – Zona D, Jalan Jenderal Sudirman Kav. 5-6 Kel. Karet Tengsin, Kec. Tanah Abang</p>
          <p class="office-phone">📞 +62 21 8665 6809 / 8665 6810</p>
        </div>
        
        <div class="office-item reveal reveal-up" data-delay="150">
          <img src="/images/lokasi pabrik.jpeg" alt="Mill Site" />
          <h3 class="office-title">Mill Site</h3>
          <p class="office-address">Desa Banuayu, Kec. Empat Petulai Dangku, Kab. Muara Enim, Sumatera Selatan</p>
          <p class="office-phone">📞 (62) (713) 324150 – 324160</p>
        </div>
        
        <div class="office-item reveal reveal-up" data-delay="300">
          <img src="/images/palembang.jpeg" alt="Palembang Office" />
          <h3 class="office-title">Palembang</h3>
          <p class="office-address">Ruko Blok I/29, Komplek PTC Mall. Jl. R. Soekamto Palembang 30114, Sumatera Selatan</p>
          <p class="office-phone">📞 (62) (711) 382409</p>
        </div>
        
        <div class="office-item reveal reveal-up" data-delay="450">
          <img src="/images/tarahan.jpeg" alt="Tarahan Office" />
          <h3 class="office-title">Tarahan</h3>
          <p class="office-address">Jl. Soekarno Hatta Km. 14, Batu Serampok Kel. Srengsem Kec. Panjang, Bandar Lampung</p>
          <p class="office-phone">📞 (62) (721) 342311, 31318</p>
        </div>
      </div>
    </section>

  
<section class="company-brief-section">
  <div class="brief-inner-wrapper">
    
    <div class="brief-left-content">
      <p>
        PT. Tanjungenim Lestari Pulp and Paper (PT. TEL), is one of the most exciting pulp mills in Indonesia today and the only pulp mill in the world to produce high-quality, bleached-hardwood Kraft pulp with 100 percent plantation grown Acacia mangium and Eucalyptus Pellita trees.
      </p>
    </div>

    <div class="brief-right-logos">
      <div class="logo-row-item">
        <img src="/images/smp.jpeg" alt="YPTEL - SMP Lematang Lestari" />
      </div>
      <div class="logo-row-item">
        <img src="/images/sd.jpeg" alt="YPTEL - SD Lematang Lestari" />
      </div>
      <div class="logo-row-item">
        <img src="/images/pgtk.jpeg" alt="YPTEL - PGTK LEMATANG LESTARI" />
      </div>
    </div>

  </div>
</section>
    <footer class="site-footer">
  <div class="footer-container">
    
    <div class="footer-left-content">
      </div>
<footer class="site-footer">
      <div class="footer-container">
        <div class="footer-left-content"></div>
        <div class="footer-copyright">
          <p>Copyright 2026 PT TELPP. All right reserved.</p>
        </div>
      </div>
    </footer>

  </div>
</footer>

  </div>
</template>

<script setup>
import { ref, computed, nextTick, onMounted, onUnmounted } from 'vue'
import { RouterLink } from 'vue-router'
import api from '../services/api'

const BASE_URL = 'http://localhost:8080'
const fallbackImg = 'https://placehold.co/600x400/e8e8e8/999?text=News'
const latestNews = ref([])

const getImageUrl = (path) => {
  if (!path) {return fallbackImg}
  if (path.startsWith('http')) {return path}
  return `${BASE_URL}/${path.replace(/^\//, '')}`
}

const formatDate = (dateStr) => {
  if (!dateStr) {return ''}
  const d = new Date(dateStr)
  return d.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
}

const fetchLatestNews = async () => {
  try {
    const res = await api.get('/news?page=1')
    const data = res.data
    const rawList = data.data || data.news || data || []
    latestNews.value = rawList
      .filter(item => item.is_published)
      .slice(0, 4)
      .map(item => ({
        id: item.ID || item.id,
        title: item.title,
        excerpt: item.summary || '',
        author: item.author || 'Admin',
        date: item.CreatedAt || item.created_at || item.date,
        image: item.Images?.[0]?.image_url || item.thumbnail_path || '',
      }))
    await nextTick()
    initScrollReveal()
  } catch (e) {
    console.error('Failed to fetch latest news:', e)
  }
}

const heroSlides = [
  {
    src: '/images/jalan-santai.jpeg',
    sub: 'Fun Walk',
    title: 'Big ',
    accent: 'Family',
    type: 'normal'
  },
  {
    src: '/images/kereta.jpeg',
    sub: 'PT TEL',
    title: 'Plan with ',
    accent: 'Ingenuity',
    type: 'normal'
  },
  {
    src: '/images/pulp2.jpeg',
    sub: 'Pulp and Paper',
    title: 'Build With ',
    accent: 'Integrity',
    type: 'normal'
  },
  {
    src: '/images/gedung.jpeg',
    sub: '',
    title: 'Continuous Improvement is Better Than Delayed Perfection',
    accent: '',
    type: 'quote'
  },
]


const sustainabilityItems = [
  {
    img: '/images/penyerahan.jpeg',
    title: 'Local Community Development',
    link: '/sustainability/csr/community' // Diperbaiki agar sesuai dengan rute router index.js Anda
  },
  {
    img: '/images/people1.jpeg',
    title: 'People Development', 
    link: '/sustainability/people-development' 
  },
  {
    img: '/images/gaming.jpeg',
    title: 'Pulp Process',
    link: '/sustainability/pulp-process' 
  },
  {
    img: '/images/people2.jpeg',
    title: 'Supply Chain Management',
    link: '/sustainability/supply-chain' 
  },
  {
    img: '/images/people3.jpeg',
    title: 'Safety & Health Office',
    link: '/sustainability/safety-health' 
  },
  {
    img: '/images/hutan3.jpeg',
    title: 'Forest Management',
    link: '/sustainability/forest-management'
  }
]

const currentSlide = ref(0)
let interval = null

// Menentukan akordeon mana yang aktif (terbuka secara bawaan adalah 'certified')
const activeAccordion = ref('certified')

// Fungsi untuk membuka dan menutup panel gambar akordeon
function toggleAccordion(panelName) {
  if (activeAccordion.value === panelName) {
    activeAccordion.value = null // Menutup jika panel yang sama diklik ulang
  } else {
    activeAccordion.value = panelName // Membuka panel baru
  }
}

// Per-character animation — setiap huruf muncul satu per satu
const currentChars = computed(() => {
  const slide = heroSlides[currentSlide.value]
  const title = (slide.title || '')
  const accent = (slide.accent || '').trim()
  const chars = []
  let animIdx = 0

  for (const ch of title) {
    chars.push({ text: ch, delay: animIdx * 45, type: 'title' })
    if (ch !== ' ') {animIdx++}
  }
  for (const ch of accent) {
    chars.push({ text: ch, delay: animIdx * 45, type: 'accent' })
    if (ch !== ' ') {animIdx++}
  }
  return chars
})

onMounted(() => {
  interval = setInterval(() => {
    currentSlide.value = (currentSlide.value + 1) % heroSlides.length
  }, 5500)
  initScrollReveal()
  fetchLatestNews()
})

onUnmounted(() => {
  clearInterval(interval)
})

function initScrollReveal() {
  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          const delay = entry.target.dataset.delay || 0
          setTimeout(() => {
            entry.target.classList.add('revealed')
          }, Number(delay))
          observer.unobserve(entry.target)
        }
      })
    },
    { threshold: 0.15, rootMargin: '0px 0px -60px 0px' }
  )
  nextTick(() => {
    document.querySelectorAll('.reveal').forEach((el) => observer.observe(el))
  })
}
</script>

<style scoped>
* {
  box-sizing: border-box;
}

.page {
  font-family: 'Inter', Arial, sans-serif;
  color: #2b2b2b;
  overflow-x: hidden;
}

.eyebrow {
  display: inline-block;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 2px;
  text-transform: uppercase;
  color: #5F9E42;
  margin-bottom: 12px;
}

.eyebrow--light {
  color: #ffffff;
}

.arrow {
  display: inline-block;
}

/* ============ HERO ============ */
.hero {
  position: relative;
  width: 100%;
  height: 480px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.hero-bg {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  opacity: 0;
  transform: scale(1.45);
  transition: opacity 1.2s ease-in-out;
}

.hero-bg--active {
  opacity: 1;
  animation: heroZoom 7s cubic-bezier(0.25, 0.46, 0.45, 0.94) forwards;
}

@media (prefers-reduced-motion: reduce) {
  .hero-bg--active {
    animation: none;
  }
}

@keyframes heroZoom {
  from { transform: scale(1.45); }
  to   { transform: scale(1.0); }
}

.hero-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, rgba(20, 30, 20, 0.35) 0%, rgba(20, 30, 20, 0.55) 100%);
}

.hero-content {
  position: relative;
  z-index: 2;
  text-align: center;
  color: #ffffff;
  padding: 0 60px;
  width: 100%;
  max-width: 950px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

/* Wrapper flex column agar hero-sub dan h1 ditumpuk vertikal di tengah */
.hero-title-block {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  width: 100%;
}

/* "Fun Walk" / "PT TEL" / "Pulp and Paper" — tepat di atas judul besar */
.hero-sub {
  display: inline-block;
  font-size: 13px;
  font-weight: 400;
  letter-spacing: 3px;
  text-transform: uppercase;
  color: rgba(255, 255, 255, 0.75);
  margin-bottom: 14px;
  padding-bottom: 10px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.4);
  text-align: center;
}

.hero-title-block h1 {
  font-size: 52px;
  font-weight: 700;
  margin: 0;
  line-height: 1.15;
  width: 100%;
  text-align: center;
}

/* ---- Animasi per huruf ---- */
.char-container {
  display: inline;
}

.char {
  display: inline-block;
  opacity: 0;
  transform: translateY(30px) scaleY(0.8);
  animation: charIn 0.5s cubic-bezier(0.22, 1, 0.36, 1) forwards;
}

.char-space {
  width: 0.28em;
  transform: none;
  animation: none;
  opacity: 1;
}

.char.accent {
  color: #8FD14F;
}

@keyframes charIn {
  to {
    opacity: 1;
    transform: translateY(0) scaleY(1);
  }
}

.hero-quote {
  font-size: 38px !important;
  font-weight: 700;
  text-align: center !important;
  text-transform: uppercase;
  letter-spacing: 2px;
  line-height: 1.25;
  width: 100%;
}

/* ---- Fade transition antar slide ---- */
.fade-title-enter-active {
  transition: opacity 0.3s ease;
}

.fade-title-leave-active {
  transition: opacity 0.4s ease;
}

.fade-title-enter-from,
.fade-title-leave-to {
  opacity: 0;
}

.scroll-down {
  position: absolute;
  bottom: -20px;
  right: 40px;
  width: 44px;
  height: 44px;
  border-radius: 50%;
  background: #1f6e2e;
  color: #ffffff;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  z-index: 3;
  font-size: 18px;
}

/* ============ ABOUT ============ */
.about {
  padding: 64px 60px 40px;
  background: #ffffff;
}

.about-inner {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 48px;
  align-items: center;
}

.about-text h2,
.about-text .eyebrow {
  margin: 0;
}

.about-text .eyebrow {
  font-size: 22px;
  font-weight: 800;
  letter-spacing: 1px;
  color: #1f1f1f;
}

.about-text p {
  font-size: 13px;
  line-height: 1.9;
  color: #555555;
  margin: 18px 0 24px;
}

.btn-light {
  background: #f3f3f3;
  border: none;
  color: #333333;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 1px;
  text-transform: uppercase;
  padding: 12px 28px;
  border-radius: 4px;
  cursor: pointer;
}

.about-media {
  position: relative;
}

.about-img {
  width: 100%;
  height: 280px;
  object-fit: cover;
  border-radius: 4px;
  display: block;
}

.about-badge {
  position: absolute;
  bottom: -20px;
  right: -20px;
  background: #2f7a3c;
  color: #ffffff;
  width: 110px;
  height: 110px;
  border-radius: 4px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
}

.badge-number {
  font-size: 26px;
  font-weight: 800;
}

.badge-label {
  font-size: 9px;
  letter-spacing: 1px;
  text-transform: uppercase;
  margin-top: 6px;
}

/* ============ SECTION HEADINGS ============ */
.section-heading {
  text-align: center;
  margin-bottom: 40px;
}

.section-heading h2 {
  font-size: 16px;
  font-weight: 800;
  letter-spacing: 3px;
  text-transform: uppercase;
  margin: 0 0 12px;
  color: #1f1f1f;
}

.section-heading--light h2 {
  color: #ffffff;
}

.heading-underline {
  display: inline-block;
  width: 50px;
  height: 2px;
  background: #5F9E42;
}

.heading-underline--light {
  background: #8FD14F;
}

/* ============ COMMUNITY ============ */
.community {
  padding: 80px 60px;
  background: #f6f7f5;
}

.community-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 32px;
  max-width: 900px;
  margin: 0 auto;
}

/* CSS perbaikan agar router-link bertingkah laku seperti div biasa tanpa merusak style teks */
.community-card {
  background: #ffffff;
  padding: 32px;
  border-radius: 4px;
  text-decoration: none;
  color: inherit;
  display: block;
  cursor: pointer;
}

.icon-box {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 18px;
}

.icon-box img {
  width: 20px;
  height: 20px;
}

.icon-pink   { background: #fde6ea; }
.icon-green { background: #e3f3e1; }
.icon-teal  { background: #e0f5f1; }
.icon-blue  { background: #e3edfb; }

.community-card h3 {
  font-size: 15px;
  font-weight: 700;
  margin: 0 0 12px;
}

.community-card p {
  font-size: 12px;
  line-height: 1.8;
  color: #777777;
  margin: 0;
}

/* ============ SUSTAINABILITY ============ */
.sustainability {
  background: #14532d;
  padding: 70px 60px;
}

.sustainability-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
  max-width: 1100px;
  margin: 0 auto;
}

.sustainability-item {
  position: relative;
  overflow: hidden;
  border-radius: 4px;
  height: 200px;
}

.sustainability-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
  transition: transform 0.4s ease;
}

.sustainability-overlay {
  position: absolute;
  inset: 0;
  background: rgba(95, 158, 66, 0.92);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.35s ease;
}

.overlay-content {
  text-align: center;
  color: #ffffff;
  padding: 16px;
  transform: translateY(15px);
  transition: transform 0.35s ease;
}

.overlay-content h3 {
  font-size: 15px;
  font-weight: 700;
  margin: 0 0 8px 0;
  letter-spacing: 0.5px;
}

.overlay-divider {
  display: block;
  width: 35px;
  height: 2px;
  background: #ffffff;
  margin: 0 auto 12px;
}

.search-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  color: #ffffff;
  text-decoration: none;
  font-size: 20px;
  transition: transform 0.2s ease;
}

.sustainability-item:hover .sustainability-overlay {
  opacity: 1;
}

.sustainability-item:hover .overlay-content {
  transform: translateY(0);
}

.sustainability-item:hover img {
  transform: scale(1.05);
}

.search-btn:hover {
  transform: scale(1.2);
}

/* ============ PRODUCT ============ */
.product {
  padding: 80px 60px;
  background: #ffffff;
}

.product-inner {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 48px;
  align-items: center;
  max-width: 1100px;
  margin: 0 auto;
}

.product-text h2 {
  font-size: 30px;
  font-weight: 800;
  line-height: 1.3;
  margin: 0 0 16px;
  color: #1f1f1f;
}

.product-sub {
  font-size: 13px;
  color: #777777;
  margin: 0 0 28px;
}

.accordion {
  display: flex;
  flex-direction: column;
  gap: 1px;
}

.accordion-group {
  display: flex;
  flex-direction: column;
  width: 100%;
}

.accordion-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #ffffff;
  border: 1px solid #e6e6e6;
  padding: 16px 20px;
  font-size: 12px;
  font-weight: 600;
  letter-spacing: 0.5px;
  color: #333333;
  cursor: pointer;
  text-align: left;
  width: 100%;
  transition: background-color 0.2s ease, color 0.2s ease;
}

/* Gaya khusus ketika baris akordeon terbuka (Header Gelap Teks Putih) */
.accordion-item--active {
  background: #1e293b !important;
  color: #ffffff !important;
}

.accordion-icon {
  color: #999999;
}

.accordion-item--active .accordion-icon {
  color: #ffffff;
}

/* Panel penampung gambar di bawah menu akordeon */
.accordion-panel {
  padding: 10px 0;
  background: #ffffff;
  border-left: 1px solid #e6e6e6;
  border-right: 1px solid #e6e6e6;
  border-bottom: 1px solid #e6e6e6;
  display: flex;
  justify-content: center;
}

.accordion-panel-img {
  width: 100%;
  max-height: 250px;
  object-fit: cover;
  display: block;
}

.product-media img {
  width: 100%;
  height: 320px;
  object-fit: cover;
  border-radius: 4px;
  display: block;
}

/* ============ NEWS ============ */
.news {
  padding: 70px 60px;
  background: #f6f7f5;
}

.news-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  max-width: 1100px;
  margin: 0 auto 36px;
}

.news-header h2 {
  font-size: 22px;
  font-weight: 800;
  margin: 0;
  color: #1f1f1f;
}

.view-all {
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 1px;
  text-transform: uppercase;
  color: #5F9E42;
  text-decoration: none;
}

.news-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 24px;
  max-width: 1100px;
  margin: 0 auto;
}

.news-card {
  display: flex;
  flex-direction: column;
  background: #ffffff;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.news-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.12);
}

.news-card .block {
  display: block;
  overflow: hidden;
}

.news-thumb {
  width: 100%;
  height: 200px;
  object-fit: cover;
  display: block;
  transition: transform 0.3s ease;
}

.news-card:hover .news-thumb {
  transform: scale(1.05);
}

.news-body {
  padding: 20px;
  flex: 1;
  display: flex;
  flex-direction: column;
}

.news-meta {
  display: flex;
  gap: 12px;
  font-size: 11px;
  color: #999999;
  margin-bottom: 10px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  font-weight: 600;
}

.news-meta .news-author {
  color: #5F9E42;
}

.news-body h3 {
  font-size: 16px;
  font-weight: 700;
  line-height: 1.4;
  margin: 0 0 10px;
  color: #1f1f1f;
}

.news-body h3:hover {
  color: #5F9E42;
}

.news-body p {
  font-size: 13px;
  line-height: 1.6;
  color: #666666;
  margin: 0;
  flex: 1;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* ============ SPEAK UP ============ */
.speakup {
  background: #1f6e2e;
  padding: 60px;
  text-align: center;
}

.btn-outline-light {
  background: #ffffff;
  border: none;
  color: #1f6e2e;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 1px;
  text-transform: uppercase;
  padding: 14px 36px;
  border-radius: 4px;
  cursor: pointer;
}

/* ============ OFFICES ============ */
/* ============ OFFICES (UPDATED) ============ */
.offices {
  padding: 80px 60px;
  background: #ffffff;
}

.offices-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 30px;
  max-width: 1200px;
  margin: 0 auto;
}

.office-item {
  display: flex;
  flex-direction: column;
  text-align: left; /* Aligns content to the left as seen in the reference image */
}

.office-item img {
  width: 100%;
  height: 180px;
  object-fit: cover;
  border-radius: 4px;
  display: block;
  margin-bottom: 20px;
}

.office-title {
  font-size: 20px;
  font-weight: 800;
  color: #1a4f8b; /* Deep blue heading color matching the image */
  text-transform: uppercase;
  margin: 0 0 12px 0;
  letter-spacing: 0.5px;
}

.office-address {
  font-size: 12px;
  line-height: 1.6;
  color: #777777;
  margin: 0 0 10px 0;
  flex-grow: 1; /* Ensures buttons/phones line up if text heights vary */
}

.office-phone {
  font-size: 12px;
  color: #666666;
  margin: 0;
  font-weight: 500;
}

@media (max-width: 900px) {
  .offices-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 500px) {
  .offices-grid {
    grid-template-columns: 1fr;
  }
}

.offices-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
  max-width: 1100px;
  margin: 0 auto;
}

.office-item {
  position: relative;
  border-radius: 4px;
  overflow: hidden;
}

.office-item img {
  width: 100%;
  height: 180px;
  object-fit: cover;
  display: block;
}

.office-label {
  position: absolute;
  bottom: 12px;
  left: 0;
  right: 0;
  text-align: center;
  color: #ffffff;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 1px;
  text-transform: uppercase;
}

/* ============ FOOTER ============ */
.site-footer {
  width: 100%;
  /* Warna hijau tua disesuaikan dengan identitas PT TEL pada gambar */
  background-color: #5F9E42; 
  padding: 25px 0; /* Memberikan ruang tinggi baris yang pas */
  margin-top: 40px; /* Jarak pemisah dari konten di atasnya */
}

.footer-container {
  display: flex;
  align-items: center;
  justify-content: center; /* DIUBAH: Membuat konten berkumpul di tengah */
  max-width: 1140px; /* Menyelaraskan lebar dengan konten atas website */
  margin: 0 auto;
  padding: 0 40px; /* Jarak aman agar teks tidak terlalu menempel ke pinggir layar */
}

/* Kolom Kiri Footer */
.footer-left-content {
  display: none; /* DIUBAH: Disembunyikan karena kosong, agar tidak memakan space */
}

/* Kolom Kanan Footer (Teks Copyright) */
.footer-copyright {
  text-align: center; /* DIUBAH: Memastikan teks rata tengah */
  width: 100%; /* Memastikan block mengambil ruang penuh untuk fungsi centering */
}

.footer-copyright p {
  font-family: Arial, sans-serif;
  font-size: 13px; /* Ukuran teks kecil, tipis, dan normal untuk footer */
  color: #ffffff;  /* Warna teks putih bersih agar kontras dengan background hijau */
  margin: 0;
  font-weight: normal;
  letter-spacing: 0.3px;
}

/* Responsif untuk Layar Handphone (Mobile) */
@media (max-width: 768px) {
  .footer-container {
    flex-direction: column;
    gap: 15px;
    padding: 0 20px;
  }
  
  .footer-copyright {
    text-align: center; /* Di HP otomatis pindah ke tengah bawah agar rapi */
    width: 100%;
  }
  
  .footer-copyright p {
    font-size: 12px; /* Ukuran sedikit disesuaikan untuk layar kecil */
  }
}
.newsletter-form {
  display: flex;
  margin-top: 16px;
  border: 1px solid #3a6b46;
  border-radius: 4px;
  overflow: hidden;
}

.newsletter-form input {
  flex: 1;
  background: transparent;
  border: none;
  padding: 10px 12px;
  color: #ffffff;
  font-size: 12px;
  outline: none;
}

.newsletter-form button {
  background: #5F9E42;
  border: none;
  color: #ffffff;
  width: 40px;
  cursor: pointer;
}

.footer-bottom {
  max-width: 1200px;
  margin: 40px auto 0;
  padding-top: 20px;
  border-top: 1px solid #2a5c39;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.footer-arrows {
  display: flex;
  gap: 8px;
}

.footer-arrows button {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  border: 1px solid #3a6b46;
  background: transparent;
  color: #ffffff;
  cursor: pointer;
}

.footer-bottom p {
  font-size: 11px;
  color: #9fb8a1;
  margin: 0;
}

/* ============ COMPANY BRIEF & LOGOS ============ */
/* ============ PENYELARASAN LOGO DAN TEKS ============ */
.company-brief-section {
  width: 100%;
  background-color: #ffffff;
  padding: 50px 0;
}

.brief-inner-wrapper {
  display: flex;
  align-items: center;      /* Membuat teks dan logo sejajar tengah secara vertikal */
  justify-content: center;   /* Membuat konten berkumpul rapi di tengah halaman */
  max-width: 1140px;
  margin: 0 auto;
  padding: 0 40px;
  gap: 60px;                 /* Jarak ideal antara teks di kiri dan logo di kanan (tidak terlalu jauh) */
}

/* Kolom Teks Kiri - Dikunci Lebarnya */
.brief-left-content {
  width: 520px;              /* Mengunci lebar teks agar pas sesuai bentuk paragraf */
  flex-shrink: 0;
}

.brief-left-content p {
  font-family: Arial, sans-serif;
  font-size: 14.5px;
  line-height: 1.8;
  color: #7c7c7c;
  margin: 0;
  text-align: justify;
}

/* Kolom Logo Kanan - Dikunci Lebarnya */
.brief-right-logos {
  width: 320px;              /* Mengunci area lebar logo sekolah agar tidak lari ke kanan */
  display: flex;
  flex-direction: column;
  gap: 14px;                 /* Jarak vertikal antar logo sekolah agar rapat & rapi */
  flex-shrink: 0;
}

.logo-row-item {
  display: flex;
  align-items: center;
  justify-content: flex-start; /* Membuat logo berjejer rapi lurus ke bawah */
  width: 100%;
}

/* KUNCI UTAMA: Memaksa Gambar 1, 2, 3 SINKRON & SAMA RATA */
.logo-row-item img {
  height: 46px;              /* Tinggi dikunci rata 46px agar ukuran logo 1, 2, 3 sama besar */
  width: 280px;              /* Lebar dikunci pasti 280px agar panjang teks sekolahnya sinkron */
  display: block;
  object-fit: contain;       /* Mencegah gambar menjadi gepeng akibat dipaksa ukuran di atas */
  object-position: left center; /* Memaksa isi logo merapat lurus dari sebelah kiri */
}

/* Responsif untuk Layar Handphone (Otomatis Rata Tengah) */
@media (max-width: 900px) {
  .brief-inner-wrapper {
    flex-direction: column;
    gap: 35px;
    padding: 0 20px;
  }
  
  .brief-left-content, 
  .brief-right-logos {
    width: 100%;
    max-width: 100%;
  }
  
  .brief-left-content p {
    text-align: center;
  }
  
  .logo-row-item {
    justify-content: center;
  }
  
  .logo-row-item img {
    height: 40px;
    width: auto;
    max-width: 260px;
    object-position: center;
  }
}






/* Responsif untuk Layar Smartphone / Kecil */
@media (max-width: 768px) {
  .brief-left-content, 
  .brief-right-logos {
    display: block;
    width: 100%;
    padding-right: 0;
  }
  
  .brief-left-content {
    margin-bottom: 30px;
  }
  
  .brief-left-content p {
    text-align: center;
  }
  
  .logo-row-item img {
    height: 32px; /* Sedikit lebih kecil jika dibuka di HP */
    margin: 0 auto; /* Membuat logo otomatis rata tengah di layar HP */
  }
}


/* Penyesuaian untuk layar tablet dan mobile */
@media (max-width: 900px) {
  .brief-container {
    flex-direction: column;
    padding: 0 24px;
    gap: 40px;
  }
  
  .brief-text, .brief-logos {
    flex: 0 0 100%;
    padding-right: 0;
  }
  
  .brief-text p {
    text-align: center;
  }

  .logo-row {
    justify-content: center;
  }
  
  .logo-row img {
    height: 46px; /* Sedikit disesuaikan untuk mobile */
  }
}

.brief-container {
  display: table;
  width: 100%;
  max-width: 960px; /* Membatasi lebar total agar tidak terlalu melar ke samping */
  margin: 0 auto;
  padding: 0 20px;
}

.brief-text {
  display: table-cell;
  width: 55%; /* Teks mengambil porsi sedikit lebih besar */
  vertical-align: middle;
  padding-right: 40px;
}

.brief-text p {
  font-family: 'Open Sans', Arial, sans-serif;
  font-size: 13.5px;
  line-height: 1.75;
  color: #8c8c8c; /* Warna abu-abu soft */
  margin: 0;
  text-align: justify;
}

.brief-logos {
  display: table-cell;
  width: 45%;
  vertical-align: middle;
  text-align: left;
}

.logo-row {
  margin-bottom: 12px; /* Jarak antar baris logo */
}

.logo-row:last-child {
  margin-bottom: 0;
}

.logo-row img {
  height: 38px; /* Mengunci tinggi gambar agar pas dan tidak kebesaran */
  width: auto;
  display: block;
  object-fit: contain;
}

/* Responsif untuk Handphone / Layar Kecil */
@media (max-width: 768px) {
  .brief-text, .brief-logos {
    display: block;
    width: 100%;
    padding-right: 0;
  }
  
  .brief-logos {
    margin-top: 30px;
  }
  
  .logo-row img {
    height: 34px; /* Sedikit lebih kecil di mobile */
    margin: 0 auto 12px auto;
  }
  
  .brief-text p {
    text-align: center;
  }
}

.brief-container {
  display: grid;
  grid-template-columns: 1.2fr 1fr; /* Membagi kolom teks sedikit lebih lebar dari kolom logo */
  gap: 60px;
  max-width: 1100px;
  margin: 0 auto;
  align-items: center;
}

.brief-text p {
  font-size: 14px;
  line-height: 1.8;
  color: #777777; /* Warna abu-abu soft sesuai gambar */
  margin: 0;
}

.brief-logos {
  display: flex;
  flex-direction: column;
  gap: 16px; /* Jarak vertikal antar 3 logo */
}

.logo-item img {
  max-width: 100%;
  height: auto;
  display: block;
  object-fit: contain;
}

/* Penyesuaian Responsif untuk perangkat Mobile */
@media (max-width: 900px) {
  .brief-container {
    grid-template-columns: 1fr;
    gap: 40px;
    text-align: center;
  }
  
  .brief-logos {
    align-items: center;
  }
}

/* ============ SCROLL REVEAL ============ */
.reveal {
  opacity: 0;
  transition: opacity 0.8s cubic-bezier(0.22, 1, 0.36, 1),
              transform 0.8s cubic-bezier(0.22, 1, 0.36, 1);
}

.reveal.revealed {
  opacity: 1;
  transform: translate(0, 0) scale(1) !important;
}

.reveal-up {
  transform: translateY(50px);
}

.reveal-down {
  transform: translateY(-50px);
}

.reveal-left {
  transform: translateX(-60px);
}

.reveal-right {
  transform: translateX(60px);
}

.reveal-scale {
  transform: scale(0.85);
}

/* ============ RESPONSIVE ============ */
@media (max-width: 900px) {
  .about-inner,
  .product-inner {
    grid-template-columns: 1fr;
  }

  .community-grid,
  .news-grid,
  .sustainability-grid {
    grid-template-columns: 1fr;
  }

  .offices-grid {
    grid-template-columns: repeat(2, 1fr);
  }

  .footer-grid {
    grid-template-columns: 1fr 1fr;
  }
}

@media (max-width: 500px) {
  .hero-content {
    padding: 0 24px;
  }

  .hero-title-block h1 {
    font-size: 36px;
  }

  .hero-quote {
    font-size: 26px !important;
  }

  .offices-grid,
  .footer-grid {
    grid-template-columns: 1fr;
  }

  .news-thumb {
    height: 160px;
  }

  .news-body h3 {
    font-size: 14px;
  }

  .about,
  .community,
  .sustainability,
  .product,
  .news,
  .offices,
  .footer {
    padding-left: 24px;
    padding-right: 24px;
  }
}
</style>