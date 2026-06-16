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
          <h1
            :key="currentSlide"
            :class="{ 'hero-quote': heroSlides[currentSlide].type === 'quote' }"
          >
            <span
              v-if="heroSlides[currentSlide].sub"
              class="hero-sub"
            >{{ heroSlides[currentSlide].sub }}</span>
            <span class="word-container">
              <span
                v-for="(w, i) in currentWords"
                :key="i"
                class="word"
                :class="{ 'accent': w.type === 'accent' }"
                :style="{ transitionDelay: w.delay + 'ms' }"
              >{{ w.text }} </span>
            </span>
          </h1>
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
            PT Tanjungenim Lestari Pulp and Paper (TELPP) is a world class
            manufacturer of high product quality and environmental friendly
            <strong>market pulp mill.</strong> This was established on June
            10, 1990, commenced construction in mid 1997 and the commercial
            operation started on May, 2000. The mill is located in 1,090 ha
            area in Niru Banuayu village, District Empat Petulai Dangku,
            Muara Enim Regency, South Sumatra province, Indonesia.
          </p>
          <button class="btn-light">Read More</button>
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
        <div class="community-card reveal reveal-up" data-delay="0">
          <div class="icon-box icon-pink">
            <img src="/images/gedung.jpeg" alt="Education icon" />
          </div>
          <h3>Education</h3>
          <p>
            Prepare a generation of quality sources from local internal
            employees as well as from the village community around the
            company of TELPP. We achieved this through collaboration with
            schools of Sumatera Selatan, Local Foundation program and
            social facilities &amp; infrastructure improvement.
          </p>
        </div>

        <div class="community-card reveal reveal-up" data-delay="150">
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
        </div>

        <div class="community-card reveal reveal-up" data-delay="300">
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
        </div>

        <div class="community-card reveal reveal-up" data-delay="450">
          <div class="icon-box icon-blue">
            <img src="/images/gedung.jpeg" alt="Health icon" />
          </div>
          <h3>Health &amp; Conservation</h3>
          <p>
            Increasing community awareness, changing the mindset of
            stakeholders and the community in health development by
            increasing community empowerment efforts through quality
            appropriate service facilities and community movements in
            healthy living.
          </p>
        </div>
      </div>
    </section>

    <section class="sustainability reveal reveal-up">
      <div class="section-heading section-heading--light">
        <h2>Sustainability</h2>
        <span class="heading-underline heading-underline--light"></span>
      </div>

      <div class="sustainability-grid">
        <div class="sustainability-item reveal reveal-up" data-delay="0">
          <img src="/images/gedung.jpeg" alt="Community team photo" />
        </div>
        <div class="sustainability-item reveal reveal-up" data-delay="200">
          <img src="/images/gedung.jpeg" alt="Plantation aerial view" />
        </div>
        <div class="sustainability-item reveal reveal-up" data-delay="400">
          <img src="/images/gedung.jpeg" alt="Field monitoring" />
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
            <button class="accordion-item">
              Certified
              <span class="accordion-icon">⌄</span>
            </button>
            <button class="accordion-item">
              Compliance Management
              <span class="accordion-icon">⌄</span>
            </button>
            <button class="accordion-item">
              The Concept
              <span class="accordion-icon">⌄</span>
            </button>
          </div>
        </div>

        <div class="product-media reveal reveal-right" data-delay="200">
          <img src="/images/gedung.jpeg" alt="Pellita pulp sheet texture" />
        </div>
      </div>
    </section>

    <section class="news reveal reveal-up">
      <div class="news-header">
        <h2>Latest news</h2>
        <a href="/news" class="view-all">More All News <span class="arrow">→</span></a>
      </div>

      <div class="news-grid">
        <article class="news-card reveal reveal-up" data-delay="0">
          <img src="/images/gedung.jpeg" alt="News thumbnail" class="news-thumb" />
          <div class="news-body">
            <span class="news-meta">
              <span class="news-author">A. Admin</span>
              <span class="news-date">Oct 11, 2022</span>
            </span>
            <h3>PT Tels Participation on the National Transportation Day Held in Bandung and Mandi...</h3>
            <p>The company decided also to commitment to environmental sustainability through achiev...</p>
          </div>
        </article>

        <article class="news-card reveal reveal-up" data-delay="150">
          <img src="/images/gedung.jpeg" alt="News thumbnail" class="news-thumb" />
          <div class="news-body">
            <span class="news-meta">
              <span class="news-author">A. Admin</span>
              <span class="news-date">Sep 16, 2022</span>
            </span>
            <h3>Karangsari Hospital Names ISPRO Physician Survival for PT. TEL</h3>
            <p>Regions of inactive on inspection of the facilities, operations and compliance...</p>
          </div>
        </article>

        <article class="news-card reveal reveal-up" data-delay="300">
          <img src="/images/gedung.jpeg" alt="News thumbnail" class="news-thumb" />
          <div class="news-body">
            <span class="news-meta">
              <span class="news-author">A. Admin</span>
              <span class="news-date">Sep 15, 2022</span>
            </span>
            <h3>PT TEL gives appreciation and award to employees who have and are approved org...</h3>
            <p>The company hosted a ceremony to honor the dedicated service of its long-term employees...</p>
          </div>
        </article>

        <article class="news-card reveal reveal-up" data-delay="450">
          <img src="/images/gedung.jpeg" alt="News thumbnail" class="news-thumb" />
          <div class="news-body">
            <span class="news-meta">
              <span class="news-author">A. Admin</span>
              <span class="news-date">Sep 15, 2022</span>
            </span>
            <h3>Telprich and PLTBG OPENS NEW THP THERESONG OUT QUALITY EQUIPMENT</h3>
            <p>The annual sports for management concluded with thriving that match outstanding team spirit and...</p>
          </div>
        </article>
      </div>
    </section>

    <section class="speakup reveal reveal-up">
      <span class="eyebrow eyebrow--light">Speak Up</span>
      <button class="btn-outline-light">Tell Us Your Story</button>
    </section>

    <section class="offices reveal reveal-up">
      <div class="section-heading">
        <h2>Our Offices</h2>
        <span class="heading-underline"></span>
      </div>

      <div class="offices-grid">
        <div class="office-item reveal reveal-up" data-delay="0">
          <img src="/images/gedung.jpeg" alt="Headquarters building" />
          <span class="office-label">Headquarters</span>
        </div>
        <div class="office-item reveal reveal-up" data-delay="150">
          <img src="/images/gedung.jpeg" alt="Admin building" />
          <span class="office-label">Admin Bldg</span>
        </div>
        <div class="office-item reveal reveal-up" data-delay="300">
          <img src="/images/gedung.jpeg" alt="Tech hub" />
          <span class="office-label">Tech Hub</span>
        </div>
        <div class="office-item reveal reveal-up" data-delay="450">
          <img src="/images/gedung.jpeg" alt="Executive suite" />
          <span class="office-label">Executive Suite</span>
        </div>
      </div>
    </section>

    <footer class="footer reveal reveal-up">
      <div class="footer-grid">
        <div class="footer-col reveal reveal-up" data-delay="0">
          <h4 class="footer-logo">TELPP</h4>
          <p>
            A leading global provider of high-quality, sustainable pulp
            products, striving toward innovation from the heart of Indonesia.
          </p>
        </div>

        <div class="footer-col reveal reveal-up" data-delay="150">
          <h5>Quick Links</h5>
          <ul>
            <li><a href="#">Privacy Policy</a></li>
            <li><a href="#">Terms of Service</a></li>
            <li><a href="#">Compliance</a></li>
            <li><a href="#">Sitemap</a></li>
            <li><a href="#">Environmental Report</a></li>
          </ul>
        </div>

        <div class="footer-col reveal reveal-up" data-delay="300">
          <h5>Our Commitment</h5>
          <p>
            Dedicated to excellence in every aspect, our operations are
            gover by world-class standards for sustainability and
            efficiency.
          </p>
        </div>

        <div class="footer-col reveal reveal-up" data-delay="450">
          <h5>Newsletter</h5>
          <p>Stay updated with our latest news and initiatives.</p>
          <form class="newsletter-form" @submit.prevent>
            <input type="email" placeholder="Enter your email" />
            <button type="submit" aria-label="Subscribe">→</button>
          </form>
        </div>
      </div>

      <div class="footer-bottom">
        <div class="footer-arrows">
          <button aria-label="Previous">‹</button>
          <button aria-label="Next">›</button>
        </div>
        <p>© 2024 PT Tanjungenim Lestari Pulp and Paper. All Rights Reserved.</p>
      </div>
    </footer>

  </div>
</template>

<script setup>
import { ref, computed, nextTick, onMounted, onUnmounted } from 'vue'

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

const currentSlide = ref(0)
let interval = null

const currentWords = computed(() => {
  const slide = heroSlides[currentSlide.value]
  const title = (slide.title || '').trim()
  const accent = (slide.accent || '').trim()
  const words = []
  if (title) {
    title.split(/\s+/).forEach((w, i) => {
      words.push({ text: w, delay: i * 100, type: 'title' })
    })
  }
  if (accent) {
    const offset = words.length
    accent.split(/\s+/).forEach((w, i) => {
      words.push({ text: w, delay: (offset + i) * 100, type: 'accent' })
    })
  }
  return words
})

onMounted(() => {
  interval = setInterval(() => {
    currentSlide.value = (currentSlide.value + 1) % heroSlides.length
  }, 4000)
  initScrollReveal()
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

/* ============ HERO (DI-PERBAIKI AGAR DI TENGAH) ============ */
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
  transform: scale(1.15);
  transition: opacity 1.2s ease-in-out;
}

.hero-bg--active {
  opacity: 1;
  animation: heroZoom 8s ease-in-out infinite alternate;
}

@keyframes heroZoom {
  from { transform: scale(1.15); }
  to   { transform: scale(1); }
}

.hero-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, rgba(20, 30, 20, 0.35) 0%, rgba(20, 30, 20, 0.55) 100%);
}

.hero-content {
  position: relative;
  z-index: 2;
  text-align: center; /* Mengubah teks judul menjadi rata tengah */
  color: #ffffff;
  padding: 0 60px;
  width: 100%;
  max-width: 950px;
  display: flex;
  flex-direction: column;
  align-items: center; /* Memaksa sub-judul dan garis pembatas ikut ke tengah */
}

.hero-sub {
  display: block;
  font-size: 13px;
  font-weight: 400;
  letter-spacing: 3px;
  text-transform: uppercase;
  color: rgba(255, 255, 255, 0.75);
  margin-bottom: 15px;
  padding-bottom: 10px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.4);
  width: fit-content;
}

.hero-content h1 {
  font-size: 52px;
  font-weight: 700;
  margin: 0;
  line-height: 1.15;
  width: 100%;
}

.word-container {
  display: inline;
}

.word {
  display: inline-block;
  opacity: 0;
  transform: translateY(20px);
  animation: wordIn 0.5s cubic-bezier(0.22, 1, 0.36, 1) forwards;
}

@keyframes wordIn {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.word.accent {
  color: #8FD14F;
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

/* ---- Fade + slide transition ---- */
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

.community-card {
  background: #ffffff;
  padding: 32px;
  border-radius: 4px;
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

.icon-pink  { background: #fde6ea; }
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
  background: #1f6e2e;
  padding: 70px 60px;
}

.sustainability-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 4px;
  max-width: 1000px;
  margin: 0 auto;
}

.sustainability-item img {
  width: 100%;
  height: 160px;
  object-fit: cover;
  display: block;
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
}

.accordion-icon {
  color: #999999;
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
  display: grid;
  grid-template-columns: 130px 1fr;
  gap: 16px;
  background: #ffffff;
  border-radius: 4px;
  overflow: hidden;
}

.news-thumb {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.news-body {
  padding: 14px 14px 14px 0;
}

.news-meta {
  display: flex;
  gap: 12px;
  font-size: 10px;
  color: #999999;
  margin-bottom: 8px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.news-body h3 {
  font-size: 13px;
  font-weight: 700;
  line-height: 1.4;
  margin: 0 0 8px;
  color: #1f1f1f;
}

.news-body p {
  font-size: 11px;
  line-height: 1.6;
  color: #888888;
  margin: 0;
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
.offices {
  padding: 80px 60px;
  background: #ffffff;
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
.footer {
  background: #14532d;
  color: #d9e6da;
  padding: 60px 60px 24px;
}

.footer-grid {
  display: grid;
  grid-template-columns: 1.4fr 1fr 1fr 1fr;
  gap: 40px;
  max-width: 1200px;
  margin: 0 auto;
}

.footer-logo {
  color: #ffffff;
  font-size: 22px;
  font-weight: 800;
  margin: 0 0 16px;
}

.footer-col h5 {
  color: #ffffff;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 1px;
  text-transform: uppercase;
  margin: 0 0 16px;
}

.footer-col p {
  font-size: 12px;
  line-height: 1.8;
  color: #b9ccba;
  margin: 0;
}

.footer-col ul {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.footer-col ul a {
  font-size: 12px;
  color: #b9ccba;
  text-decoration: none;
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

  .hero-content h1 {
    font-size: 36px;
  }

  .hero-quote {
    font-size: 26px !important;
  }

  .offices-grid,
  .footer-grid {
    grid-template-columns: 1fr;
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