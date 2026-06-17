<template>
  <div ref="container" class="relative w-full h-[650px] overflow-hidden rounded-xl" style="background:#0a1208">
    <div class="absolute top-0 left-0 right-0 z-10 pointer-events-none"
      style="padding:12px 20px;display:flex;align-items:center;justify-content:space-between;background:linear-gradient(to bottom,rgba(0,0,0,.7) 0%,transparent 100%)">
      <div>
        <h1 class="text-sm font-extrabold text-white tracking-tight">PT. Tanjungenim Lestari Pulp and Paper</h1>
        <p class="text-[9px] text-white/35 mt-0.5 tracking-widest uppercase">Interactive 3D Site Map</p>
      </div>
    </div>

    <div class="absolute top-3 right-5 z-10 flex gap-2">
      <button @click="toggleAutoRotate"
        class="rounded-full px-4 py-1.5 text-[11px] font-semibold tracking-wide transition-all cursor-pointer border"
        :class="autoRotate
          ? 'bg-green-500/15 border-green-500/40 text-green-400'
          : 'bg-white/7 border-white/13 text-white/60 hover:bg-white/14 hover:text-white'">
        &#10227; Auto Rotate
      </button>
      <button @click="topView"
        class="rounded-full px-4 py-1.5 text-[11px] font-semibold tracking-wide transition-all cursor-pointer border bg-white/7 border-white/13 text-white/60 hover:bg-white/14 hover:text-white">
        Top View
      </button>
      <button @click="resetView"
        class="rounded-full px-4 py-1.5 text-[11px] font-semibold tracking-wide transition-all cursor-pointer border bg-white/7 border-white/13 text-white/60 hover:bg-white/14 hover:text-white">
        Reset View
      </button>
    </div>

    <!-- Stage Slider -->
    <div class="absolute bottom-14 left-1/2 -translate-x-1/2 z-10 pointer-events-auto">
      <div class="bg-black/70 backdrop-blur-md rounded-2xl px-5 py-3 border border-white/10 flex items-center gap-1">
        <template v-for="(stage, i) in stages" :key="i">
          <button @click="goToStage(i)"
            class="flex flex-col items-center gap-1 px-2 py-1 rounded-lg transition-all duration-300 cursor-pointer group"
            :class="activeStage === i ? 'bg-white/10' : 'hover:bg-white/5'">
            <div class="w-7 h-7 rounded-full flex items-center justify-center text-[10px] font-bold transition-all duration-300 border-2"
              :class="activeStage === i
                ? 'border-green-400 bg-green-500/20 text-green-400 scale-110'
                : 'border-white/20 bg-white/5 text-white/50 group-hover:border-white/40 group-hover:text-white/80'">
              {{ i + 1 }}
            </div>
            <span class="text-[9px] font-semibold tracking-wide whitespace-nowrap transition-colors duration-300"
              :class="activeStage === i ? 'text-green-400' : 'text-white/40 group-hover:text-white/70'">
              {{ stage.short }}
            </span>
          </button>
          <div v-if="i < stages.length - 1" class="w-4 h-px mt-[-10px]"
            :class="activeStage > i ? 'bg-green-500/50' : 'bg-white/15'"></div>
        </template>
      </div>
    </div>

    <!-- Stage Info Panel (bottom-right, above hint) -->
    <div v-if="activeStage !== null"
      class="absolute bottom-20 right-5 z-10 pointer-events-none transition-all duration-500"
      :class="stagePanelVisible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-3'">
      <div class="bg-black/80 backdrop-blur-md rounded-xl px-4 py-3 border border-white/10 max-w-[220px]">
        <div class="flex items-center gap-2 mb-1.5">
          <span class="w-5 h-5 rounded-full flex items-center justify-center text-[9px] font-bold border"
            :style="{ borderColor: stages[activeStage].color, color: stages[activeStage].color, background: stages[activeStage].color + '20' }">
            {{ activeStage + 1 }}
          </span>
          <span class="text-xs font-bold text-white">{{ stages[activeStage].name }}</span>
        </div>
        <p class="text-[10px] text-white/50 leading-relaxed">{{ stages[activeStage].desc }}</p>
      </div>
    </div>

    <div class="absolute bottom-5 left-5 z-10 pointer-events-none flex flex-col gap-1.5">
      <div v-for="item in legendItems" :key="item.label" class="flex items-center gap-2 text-[11px] font-semibold text-white/65 tracking-wide">
        <div class="w-2.5 h-2.5 rounded-sm shrink-0" :style="{ background: item.color }"></div>
        {{ item.label }}
      </div>
    </div>

    <div class="absolute bottom-5 right-5 z-10 text-right text-[10px] text-white/25 leading-loose pointer-events-none">
      Drag to rotate &nbsp;&middot;&nbsp; Scroll to zoom<br>Click area for info
    </div>

    <div v-if="tooltip.visible"
      class="absolute z-20 pointer-events-none rounded-xl px-4 py-3 max-w-[230px] backdrop-blur-xl border border-white/10"
      :style="{ left: tooltip.x + 'px', top: tooltip.y + 'px', background: 'rgba(6,10,4,.96)' }">
      <div class="text-[9px] font-bold tracking-widest uppercase mb-1" :style="{ color: tooltip.zoneColor }">{{ tooltip.zone }}</div>
      <div class="text-sm font-extrabold text-white mb-1 leading-snug">{{ tooltip.name }}</div>
      <div class="text-xs text-white/50 leading-relaxed">{{ tooltip.desc }}</div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, onUnmounted, nextTick } from 'vue'
import * as THREE from 'three'

const container = ref(null)
const autoRotate = ref(true)
const activeStage = ref(null)
const stagePanelVisible = ref(false)

const stages = [
  { name: 'Wood & Chip Handling', short: 'Wood', desc: 'Log storage ponds, debarking drum, disk chipper, and chip screening — uniform 25mm chips ready for cooking.', color: '#d97706', theta: -0.6, phi: 0.75, radius: 55, lookAt: [-20, 0, -18] },
  { name: 'Fiber Line', short: 'Fiber', desc: 'Kamyr continuous digester (170\u00b0C), blow tank, multi-stage brownstock washers, and O\u2082 delignification.', color: '#f97316', theta: -0.15, phi: 0.82, radius: 50, lookAt: [5, 0, -8] },
  { name: 'Bleaching Plant', short: 'Bleach', desc: 'ECF bleaching sequence D\u2080/EOP/D\u2081/D\u2082 using ClO\u2082 and H\u2082O\u2082. Target brightness: 89+ % ISO.', color: '#0d9488', theta: 0.15, phi: 0.85, radius: 48, lookAt: [14, 0, -2] },
  { name: 'Pulp Machine', short: 'PM', desc: 'Fourdrinier forming machine — pulp slurry formed, pressed, and steam-dried. Capacity: 450,000 ADt/year.', color: '#0284c7', theta: 0.1, phi: 0.95, radius: 45, lookAt: [-4, 0, 4] },
  { name: 'Warehouse', short: 'Warehouse', desc: 'Air-conditioned warehouse storing finished pulp bales (250 kg each) for export via conveyor to jetty.', color: '#8b5cf6', theta: 0.3, phi: 0.9, radius: 48, lookAt: [18, 0, 4] },
  { name: 'Chemical Recovery', short: 'Recovery', desc: 'Recovery Boiler (2,400 TDS/day), evaporators, lime kiln, recausticizing, and power boiler for energy.', color: '#dc2626', theta: 0.45, phi: 0.78, radius: 55, lookAt: [12, 0, -16] }
]

const legendItems = [
  { color: '#d97706', label: 'Wood & Chip Handling' },
  { color: '#0284c7', label: 'Main Mill Process' },
  { color: '#dc2626', label: 'Chemical Recovery' },
  { color: '#0d9488', label: 'Wastewater Treatment' },
  { color: '#7c3aed', label: 'Utilities & Services' }
]

const tooltip = reactive({ visible: false, x: 0, y: 0, zone: '', name: '', desc: '', zoneColor: '' })

let scene, camera, renderer, animationId
let orTheta = -0.4, orPhi = 0.92, orRadius = 95
let tTheta = orTheta, tPhi = orPhi, tRadius = orRadius
let isDrag = false, prevX = 0, prevY = 0
let dragStartX = 0, dragStartY = 0
let currentLookAt = null
const clickables = []
const flowParticles = []

function M(col, emissive = 0, rough = 0.75, metal = 0.1) {
  return new THREE.MeshStandardMaterial({ color: col, emissive, roughness: rough, metalness: metal })
}

const MAT = {
  ground:    M(0x1a2e12, 0, 1.0),
  forest:    M(0x1a4a1a, 0x0a2a0a, 0.9),
  forest2:   M(0x2d6b2d, 0, 0.9),
  road:      M(0x2a2e24, 0, 1.0),
  sand:      M(0x8b7355, 0, 0.95),
  chip:      M(0x8B5E3C, 0x3a2010, 0.85),
  bark:      M(0x5c3d1e, 0x2a1a08, 0.9),
  logWood:   M(0x6b4423, 0, 0.85),
  steel:     M(0x4a5568, 0, 0.4, 0.8),
  concrete:  M(0x4a5060, 0, 0.9),
  white:     M(0xe0e8f0, 0, 0.6),
  lightBlue: M(0x7bbfd4, 0x0a3040, 0.3, 0.5),
  green:     M(0x0a8a3a, 0, 0.7),
  water:     M(0x1a6090, 0x082840, 0.1, 0.1),
  waterGreen:M(0x1a7050, 0x083028, 0.2),
  pool:      M(0x30a8d8, 0x083050, 0.1, 0.2),
  mill:      M(0x3a5a7a, 0x0a1a2a, 0.6, 0.3),
  boiler:    M(0xacb4bc, 0x14181c, 0.5, 0.35),
  digester:  M(0xc8651f, 0x401a05, 0.55, 0.3),
  chipbin:   M(0x6a9a7a, 0x142a1c, 0.5, 0.3),
  yellow:    M(0xc8960a, 0x503800, 0.5, 0.4),
  orange:    M(0xc84a10, 0x601808, 0.6, 0.3),
  purple:    M(0xc4c8c4, 0x1c1c1c, 0.5, 0.3),
  roofRed:   M(0xa03020, 0, 0.7),
  cream:     M(0xddd6bf, 0x100e08, 0.55, 0.15),
  silver:    M(0xb6bec4, 0x16191c, 0.45, 0.4),
}

function addClickable(mesh, zone, name, desc, zoneColor) {
  mesh.userData = { zone, name, desc, zoneColor }
  clickables.push(mesh)
}

function box(w, h, d, mat, cx, cy, cz, ry = 0) {
  const m = new THREE.Mesh(new THREE.BoxGeometry(w, h, d), mat)
  m.position.set(cx, cy, cz)
  m.rotation.y = ry
  m.castShadow = true; m.receiveShadow = true
  scene.add(m); return m
}

function cyl(rt, rb, h, seg, mat, cx, cy, cz) {
  const m = new THREE.Mesh(new THREE.CylinderGeometry(rt, rb, h, seg), mat)
  m.position.set(cx, cy, cz)
  m.castShadow = true; scene.add(m); return m
}

function cone(r, h, seg, mat, cx, cy, cz) {
  const m = new THREE.Mesh(new THREE.ConeGeometry(r, h, seg), mat)
  m.position.set(cx, cy, cz)
  m.castShadow = true; scene.add(m); return m
}

function sphere(r, mat, cx, cy, cz) {
  const m = new THREE.Mesh(new THREE.SphereGeometry(r, 12, 8), mat)
  m.position.set(cx, cy, cz)
  scene.add(m); return m
}

function treeCluster(cx, cz, count, spread, minH = 2, maxH = 5) {
  for (let i = 0; i < count; i++) {
    const ox = (Math.random() - 0.5) * spread
    const oz = (Math.random() - 0.5) * spread
    const h = minH + Math.random() * (maxH - minH)
    const tmat = Math.random() > 0.5 ? MAT.forest : MAT.forest2
    cyl(0.12, 0.14, h * 0.4, 6, M(0x5c3d11), cx + ox, h * 0.2, cz + oz)
    cone(0.3 + Math.random() * 0.4, h * 0.7, 7, tmat, cx + ox, h * 0.55 + h * 0.2, cz + oz)
  }
}

function makeLabel(text, hexCol, scale = 1) {
  const c = document.createElement('canvas')
  c.width = 320; c.height = 64
  const ctx = c.getContext('2d')
  ctx.fillStyle = hexCol
  ctx.font = `bold ${18 * scale}px Inter,Arial`
  ctx.textAlign = 'center'
  ctx.textBaseline = 'middle'
  ctx.shadowColor = 'rgba(0,0,0,.9)'
  ctx.shadowBlur = 10
  const words = text.split(' ')
  let lines = [], line = ''
  for (const w of words) {
    const t = line ? line + ' ' + w : w
    if (ctx.measureText(t).width > 290 * scale) { lines.push(line); line = w } else line = t
  }
  if (line) lines.push(line)
  const lh = 20 * scale, sy = 32 - (lines.length - 1) * lh / 2
  lines.forEach((l, i) => ctx.fillText(l, 160, sy + i * lh))
  const tex = new THREE.CanvasTexture(c)
  const sm = new THREE.SpriteMaterial({ map: tex, transparent: true, depthTest: false })
  const s = new THREE.Sprite(sm)
  s.scale.set(9 * scale, 2 * scale, 1)
  return s
}

function buildScene() {
  scene = new THREE.Scene()
  scene.background = new THREE.Color(0x0a1610)
  scene.fog = new THREE.FogExp2(0x0a1610, 0.003)

  camera = new THREE.PerspectiveCamera(40, container.value.clientWidth / container.value.clientHeight, 0.1, 800)

  renderer = new THREE.WebGLRenderer({ antialias: true })
  renderer.setSize(container.value.clientWidth, container.value.clientHeight)
  renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2))
  renderer.shadowMap.enabled = true
  renderer.shadowMap.type = THREE.PCFSoftShadowMap
  renderer.toneMapping = THREE.ACESFilmicToneMapping
  renderer.toneMappingExposure = 1.1
  container.value.appendChild(renderer.domElement)

  scene.add(new THREE.AmbientLight(0x3d5a3d, 0.8))
  const sun = new THREE.DirectionalLight(0xfff4d6, 2.0)
  sun.position.set(50, 90, 40)
  sun.castShadow = true
  sun.shadow.mapSize.set(4096, 4096)
  sun.shadow.camera.left = -150; sun.shadow.camera.right = 150
  sun.shadow.camera.top = 150; sun.shadow.camera.bottom = -150
  sun.shadow.camera.far = 250
  sun.shadow.bias = -0.001
  scene.add(sun)
  const fl1 = new THREE.DirectionalLight(0x7ab0cc, 0.4)
  fl1.position.set(-40, 25, -30)
  scene.add(fl1)
  const fl2 = new THREE.DirectionalLight(0x5566bb, 0.25)
  fl2.position.set(10, 15, 50)
  scene.add(fl2)
  const rimLight = new THREE.DirectionalLight(0xffeedd, 0.3)
  rimLight.position.set(-20, 5, 40)
  scene.add(rimLight)

  const ground = new THREE.Mesh(new THREE.PlaneGeometry(300, 280), M(0x1e3318, 0, 1.0))
  ground.rotation.x = -Math.PI / 2; ground.receiveShadow = true
  scene.add(ground)

  // Factory floor pad
  const factoryPad = new THREE.Mesh(new THREE.PlaneGeometry(140, 100), M(0x3a3e36, 0, 0.95))
  factoryPad.rotation.x = -Math.PI / 2
  factoryPad.position.set(0, 0.01, -2)
  factoryPad.receiveShadow = true
  scene.add(factoryPad)

  box(90, 0.05, 4, MAT.road, 0, 0.03, 2)
  box(4, 0.05, 70, MAT.road, -28, 0.03, 0)
  box(4, 0.05, 60, MAT.road, 22, 0.03, -2)
  box(30, 0.05, 3, MAT.road, -5, 0.03, -18)
  box(30, 0.05, 3, MAT.road, 5, 0.03, 20)
  box(3, 0.05, 40, MAT.road, 8, 0.03, 8)
  for (let i = -40; i < 44; i += 3.5) box(1.5, 0.06, 0.15, MAT.white, i, 0.03, 2)

  treeCluster(-60, 0, 50, 30, 3, 7)
  treeCluster(55, 0, 40, 25, 3, 7)
  treeCluster(0, -65, 50, 35, 3, 7)
  treeCluster(0, 65, 40, 30, 3, 7)
  treeCluster(-20, -10, 20, 10, 2, 5)
  treeCluster(5, -25, 18, 8, 2, 5)
  treeCluster(30, 5, 25, 12, 3, 6)
  treeCluster(-10, 30, 15, 8, 2, 4)
  treeCluster(15, 30, 20, 10, 2, 5)
  treeCluster(-35, 40, 20, 10, 3, 6)
  treeCluster(40, -20, 20, 10, 3, 6)
  treeCluster(-5, 55, 30, 12, 3, 6)

  // ZONE 1: LOG STORAGE PONDS
  {
    const base1 = box(16, 0.1, 6, MAT.sand, -30, 0.05, -28)
    addClickable(base1, 'Wood Handling', 'Log Storage Pond A', 'Storage pond for freshly harvested logs awaiting debarking. Logs are kept wet to preserve fiber quality.', '#d97706')
    for (let r = 0; r < 3; r++) for (let c = 0; c < 5; c++) box(2.5, 0.8, 1.4, MAT.logWood, -37 + c * 3, 0.5 + r * 0.05, -30 + r * 1.8)
    const base2 = box(16, 0.1, 6, MAT.sand, -30, 0.05, -18)
    addClickable(base2, 'Wood Handling', 'Log Storage Pond B', 'Second log storage bay — capacity buffer ensuring continuous supply to the wood preparation plant.', '#d97706')
    for (let r = 0; r < 3; r++) for (let c = 0; c < 5; c++) box(2.5, 0.8, 1.4, MAT.logWood, -37 + c * 3, 0.5 + r * 0.05, -20 + r * 1.8)
    box(12, 0.3, 0.7, MAT.steel, -18, 0.8, -22, -0.25)
    box(12, 0.3, 0.7, MAT.steel, -18, 0.8, -27, -0.25)
  }

  // ZONE 2: WOOD HANDLING / CHIP PREPARATION
  {
    const drum = new THREE.Mesh(new THREE.CylinderGeometry(1.1, 1.1, 5, 16), MAT.mill)
    drum.rotation.z = Math.PI / 2; drum.position.set(-10, 1.3, -20)
    drum.castShadow = true; scene.add(drum)
    addClickable(drum, 'Wood Handling', 'Debarking Drum', 'Rotary drum debarker strips bark from logs. Bark is collected as biomass fuel for the Power Boiler.', '#d97706')
    const chipper = box(4, 3, 4, MAT.concrete, -5, 1.5, -22)
    addClickable(chipper, 'Wood Handling', 'Wood Chipping Plant', 'Disk chippers reduce debarked logs into uniform 25mm wood chips suitable for cooking in the digester.', '#d97706')
    box(1.5, 3.5, 1.5, M(0x374151), -3.5, 1.75, -22)
    box(20, 0.3, 0.8, MAT.steel, -10, 4, -20)
    box(0.3, 4.5, 0.3, MAT.steel, -10, 2.25, -20.4)
    box(0.3, 4.5, 0.3, MAT.steel, -10, 2.25, -19.6)
    box(0.3, 4.5, 0.3, MAT.steel, 0, 2.25, -20.4)
    box(0.3, 4.5, 0.3, MAT.steel, 0, 2.25, -19.6)
    const screener = box(5, 2, 3, MAT.concrete, 2, 1, -22)
    addClickable(screener, 'Wood Handling', 'Chip Screening', 'Vibrating screen separates oversized chips (rechipped) and fines (used as fuel) — accepts uniform chips to Chip Bin.', '#d97706')
  }

  // ZONE 3: MAIN MILL
  {
    const chipbin = cyl(1.5, 1.6, 5, 16, MAT.chipbin, -2, 2.5, -14)
    addClickable(chipbin, 'Fiber Line', 'Chip Bin', 'Pressurized chip bin with chip meter — controls feed rate of chips into the continuous digester.', '#0284c7')
    const dig1 = cyl(1.4, 1.6, 9, 20, MAT.digester, 2, 4.5, -10)
    addClickable(dig1, 'Fiber Line', 'Continuous Digester', 'Kamyr-type continuous digester — chips cooked with White Liquor (NaOH + Na\u2082S) at 170\u00b0C, dissolving lignin to produce brown stock pulp. Capacity: ~1,500 t pulp/day.', '#0284c7')
    cyl(0.9, 1.0, 1.5, 16, M(0x6a8aaa), 2, 9.75, -10)
    cyl(0.4, 0.5, 3.5, 8, MAT.steel, 2, 12, -10)
    const dig2 = cyl(1.0, 1.1, 7, 16, MAT.digester, 5, 3.5, -10)
    addClickable(dig2, 'Fiber Line', 'Impregnation Vessel', 'Pre-steaming and liquor impregnation of chips before entering main digester — ensures uniform cooking.', '#0284c7')
    const blow = cyl(1.8, 2.0, 4, 16, M(0xb05418, 0x351005, 0.55, 0.3), 2, 2, -5)
    addClickable(blow, 'Fiber Line', 'Blow Tank', 'Receives cooked pulp from digester via blow valve at atmospheric pressure — first stage of stock washing.', '#0284c7')
    const wash = box(8, 2.5, 4, MAT.mill, 8, 1.25, -8)
    addClickable(wash, 'Fiber Line', 'Washing & Screening Plant', 'Multi-stage brown stock washers and pressure screens remove black liquor and shives from cooked pulp.', '#0284c7')
    const o2a = cyl(0.9, 1.0, 6, 12, M(0x2a8aa0, 0x0a2a30, 0.5, 0.3), 14, 3, -8)
    addClickable(o2a, 'Fiber Line', 'O\u2082 Delignification', 'Oxygen delignification reduces kappa number 40-50% using O\u2082 and NaOH — reduces chlorine chemical demand in bleaching.', '#0284c7')
    cyl(0.9, 1.0, 6, 12, M(0x2a8aa0, 0x0a2a30, 0.5, 0.3), 16, 3, -8)
    const blch = box(10, 3, 5, M(0x2a6a7a), 14, 1.5, -2)
    addClickable(blch, 'Fiber Line', 'Bleaching Plant (ECF)', 'Elemental Chlorine Free (ECF) bleaching sequence D\u2080/EOP/D\u2081/D\u2082 using ClO\u2082 and H\u2082O\u2082. Target brightness: 89+ % ISO.', '#0284c7')
    const bleachCols = [M(0xcfa218, 0x4a3a08, 0.5, 0.3), M(0xcfa218, 0x4a3a08, 0.5, 0.3), MAT.cream, MAT.cream]
    for (let i = 0; i < 4; i++) cyl(0.8, 0.85, 5, 12, bleachCols[i], 10 + i * 1.8, 3.5, -2)
    const hdt = cyl(2.0, 2.2, 7, 16, MAT.cream, 14, 3.5, 5)
    addClickable(hdt, 'Fiber Line', 'HD Tower (High Density)', 'High-density storage tower — provides buffer between bleaching and pulp machine, ensures continuous operation.', '#0284c7')
    cyl(1.2, 1.5, 1.0, 16, MAT.cream, 14, 7.5, 5)
    cyl(1.1, 1.2, 7, 16, MAT.cream, -25, 3.5, 3)
    cyl(0.9, 1.0, 8, 16, MAT.cream, -23, 4, 5)
    cyl(0.8, 0.9, 5, 16, MAT.silver, -21, 2.5, 2)
    cyl(0.6, 0.65, 4, 12, MAT.silver, -19.5, 2, 4)
    box(2, 2.2, 2, M(0xb8bcc0), -15, 1.1, 3)
    box(2.2, 0.5, 2.2, M(0xb8467a, 0x300012, 0.5), -15, 2.35, 3)
    box(2, 2.2, 2, M(0xb8bcc0), -12.5, 1.1, 5)
    box(2.2, 0.5, 2.2, M(0xb8467a, 0x300012, 0.5), -12.5, 2.35, 5)
    const pmBase = box(28, 2.5, 6, M(0x35608c, 0x0a1828, 0.55, 0.3), -4, 1.25, 4)
    addClickable(pmBase, 'Fiber Line', 'Pulp Machine (PM1)', 'Fourdrinier forming machine — diluted pulp slurry formed into sheet on wire, pressed and steam-dried into market pulp bales. Design capacity: 450,000 ADt/year.', '#0284c7')
    for (let i = -12.5; i <= 12.5; i += 2.5) box(2, 0.5, 0.4, MAT.cream, i, 2.6, 4)
    for (let i = -12; i <= 12; i += 2.5) cyl(0.5, 0.5, 5.8, 8, MAT.steel, i, 1.5, 4)
    box(8, 3.2, 5.8, M(0x4a72a0, 0x0a1828, 0.5, 0.3), -12, 1.6, 4)
    const wh = box(14, 3, 7, M(0xc0c8d8), 18, 1.5, 4)
    addClickable(wh, 'Fiber Line', 'Finished Pulp Warehouse', 'Air-conditioned warehouse storing market pulp bales (250 kg each) before loading to conveyor belt to ship jetty for export.', '#0284c7')
    for (let r = 0; r < 3; r++) for (let c = 0; c < 5; c++) box(1.8, 0.9, 1.2, M(0xf0f4ff), 12 + c * 2.0, 1.55 + r * 0.95, 2 + r * 0.05)
    cyl(0.35, 0.4, 12, 8, MAT.steel, 8, -0.5, -12)
    cyl(0.35, 0.4, 10, 8, MAT.steel, 10, -0.5, -10)
    sphere(0.7, M(0xbbbbbb, 0, 0.95), 8, 11.5, -12)
    sphere(0.9, M(0xaaaaaa, 0, 0.95), 10, 10.5, -10)
  }

  // ZONE 4: CHEMICAL RECOVERY
  {
    const rb = box(5, 9, 4, MAT.boiler, 8, 4.5, -18)
    addClickable(rb, 'Chemical Recovery', 'Recovery Boiler', 'Burns concentrated black liquor (70% DS) to recover inorganic smelt and generate high-pressure steam. Design: 2,400 TDS/day.', '#dc2626')
    box(4.8, 0.5, 3.8, M(0xc8a020), 8, 9.2, -18)
    cyl(0.55, 0.65, 8, 8, MAT.steel, 8, 14, -18)
    sphere(1.0, M(0x999999, 0, 0.95), 8, 18.5, -18)
    const pb = box(4, 8, 3.5, MAT.boiler, 15, 4, -18)
    addClickable(pb, 'Chemical Recovery', 'Power Boiler (Biomass)', 'Biomass-fired power boiler burns bark and wood waste — generates steam for electricity production.', '#dc2626')
    cyl(0.45, 0.55, 7, 8, MAT.steel, 15, 12.5, -18)
    sphere(0.8, M(0x999999, 0, 0.95), 15, 16, -18)
    for (let i = 0; i < 5; i++) {
      const ev = cyl(0.65, 0.7, 5 + i * 0.3, 12, MAT.silver, 18 + i * 1.4, 2.5 + i * 0.15, -14)
      if (i === 2) addClickable(ev, 'Chemical Recovery', 'Evaporator Plant (7-effect)', 'Multiple-effect evaporator concentrates weak black liquor from 15% to 72% dissolved solids for firing in Recovery Boiler.', '#dc2626')
    }
    box(6, 0.4, 0.4, MAT.steel, 21, 5.5, -14)
    const lkDrum = new THREE.Mesh(new THREE.CylinderGeometry(0.9, 0.9, 9, 16), M(0x8a3a2a))
    lkDrum.rotation.z = Math.PI / 2; lkDrum.position.set(17, 1.1, -22)
    lkDrum.castShadow = true; scene.add(lkDrum)
    addClickable(lkDrum, 'Chemical Recovery', 'Lime Kiln', 'Rotary kiln (~150m) reburns CaCO\u2083 lime mud at 1000\u00b0C to regenerate CaO quick lime for recausticizing cycle.', '#dc2626')
    box(2, 3, 2, M(0x5a2a1a), 13, 1.5, -22)
    cyl(0.4, 0.4, 4, 8, MAT.steel, 13, 4, -22)
    cyl(0.8, 0.9, 6, 16, MAT.cream, 21, 3, -22)
    cyl(1.4, 1.4, 5, 16, MAT.cream, 24, 2.5, -25)
    const rct = cyl(1.5, 1.6, 3.5, 16, M(0x8a4a2a), 8, -0.25, -22)
    addClickable(rct, 'Chemical Recovery', 'Recausticizing Plant', 'Green liquor (Na\u2082CO\u2083+Na\u2082S) reacted with Ca(OH)\u2082 to regenerate White Liquor (NaOH+Na\u2082S) for cooking cycle.', '#dc2626')
    cyl(1.5, 1.6, 3.5, 16, M(0x8a4a2a), 5, -0.25, -22)
    const tg = box(4, 2.5, 2.5, MAT.purple, 15, -0.25, -10)
    addClickable(tg, 'Chemical Recovery', 'Steam Turbine & Generator', 'Condensing-extraction steam turbine converts boiler steam to electricity — meets mill power demand of ~80 MW.', '#dc2626')
    box(2, 2, 1.5, M(0x8a3030), 17, -0.5, -10)
    box(0.3, 0.3, 12, MAT.steel, 12, 5.5, -14)
    box(10, 0.3, 0.3, MAT.steel, 13, 2, -22)
  }

  // ZONE 5: CHIP / LOG YARD
  {
    const yardPad = box(32, 0.05, 26, MAT.sand, -30, 0.02, 22)
    addClickable(yardPad, 'Wood Handling', 'Wood Chip & Bark Yard', 'Large outdoor storage yard — holds wood chips, bark, and log reserve. Capacity ~30 days mill consumption.', '#d97706')
    const pileColors = [MAT.chip, MAT.bark, M(0x7a5030)]
    for (let col = 0; col < 3; col++) for (let row = 0; row < 4; row++) {
      const px = -40 + col * 10, pz = 12 + row * 7
      const pmat = pileColors[(col + row) % 3]
      const pile = new THREE.Mesh(new THREE.SphereGeometry(3.5, 12, 8), pmat)
      pile.scale.set(1.2, 0.38, 0.9)
      pile.position.set(px, 1.3, pz)
      pile.castShadow = true
      scene.add(pile)
      box(0.3, 1.5, 6, MAT.concrete, px + 4.5, 0.75, pz)
    }
    for (let i = 0; i < 5; i++) box(30, 0.1, 0.4, MAT.road, -30, 0.05, 9 + i * 7)
    for (let i = 0; i < 4; i++) box(0.4, 0.1, 26, MAT.road, -44 + i * 10, 0.05, 22)
    box(1, 3, 0.5, MAT.yellow, -24, 1.5, 22)
    box(8, 0.4, 0.4, MAT.yellow, -26, 3.2, 22, -0.3)
  }

  // ZONE 6: WASTEWATER TREATMENT
  {
    const aePond = box(14, 0.2, 8, MAT.waterGreen, 28, 0.1, 20)
    aePond.receiveShadow = true
    addClickable(aePond, 'WWTP', 'Aeration Pond', 'Biological treatment pond — aerobic bacteria degrade organic matter from process effluent before final clarification.', '#0d9488')
    for (let i = 0; i < 3; i++) {
      sphere(0.6, M(0x0a9060, 0x052a18, 0.3), 22 + i * 5, 0.5, 20)
      cyl(0.05, 0.05, 0.8, 6, MAT.steel, 22 + i * 5, 0.5, 20)
    }
    cyl(4, 4, 0.3, 32, MAT.water, 33, 0.15, 28)
    cyl(0.2, 0.2, 1, 8, MAT.steel, 33, 0.65, 28)
    addClickable(cyl(3.5, 3.5, 0.1, 32, M(0x1a5a80, 0x0a2a40, 0.2), 33, 0.26, 28), 'WWTP', 'Primary Clarifier 1', 'Circular settling tank — removes suspended solids from effluent via gravity settling. Sludge scraped to center.', '#0d9488')
    cyl(3.5, 3.5, 0.3, 32, MAT.water, 38, 0.15, 23)
    cyl(3.0, 3.0, 0.1, 32, M(0x1a5a80, 0x0a2a40, 0.2), 38, 0.26, 23)
    cyl(0.2, 0.2, 1, 8, MAT.steel, 38, 0.65, 23)
    addClickable(box(1, 1, 1, MAT.water, 38, 0.5, 23), 'WWTP', 'Secondary Clarifier', 'Final clarification stage before treated effluent discharge — meets environmental quality standards.', '#0d9488')
    cyl(2.5, 2.5, 0.25, 24, M(0x4a6040), 28, 0.12, 28)
    addClickable(cyl(2.0, 2.0, 0.1, 24, M(0x3a5030), 28, 0.24, 28), 'WWTP', 'Sludge Thickener', 'Concentrates biological sludge before dewatering and disposal.', '#0d9488')
    box(3, 2, 3, MAT.concrete, 28, 1, 34)
    box(3, 0.5, 3, MAT.roofRed, 28, 2.25, 34)
    box(0.2, 0.2, 10, MAT.steel, 26, 0.5, 25)
  }

  // ZONE 7: UTILITIES
  {
    const ct1 = cyl(2.5, 3.0, 4, 16, MAT.concrete, 20, -0.5, -5)
    addClickable(ct1, 'Utilities', 'Cooling Tower', 'Induced-draft cooling tower — cools process water from condensers and heat exchangers for reuse.', '#7c3aed')
    cyl(2.0, 2.5, 0.5, 16, M(0x3a3a3a), 20, 2, -5)
    sphere(2.2, M(0xdddddd, 0, 0.95), 20, 3, -5)
    const ph = box(4, 2, 3, MAT.concrete, -2, 1, -32)
    addClickable(ph, 'Utilities', 'Raw Water Intake', 'Pumps raw water from Lematang River — treated in water treatment plant to supply process and boiler feed water.', '#7c3aed')
    box(4, 0.5, 3, MAT.roofRed, -2, 2.25, -32)
    const wtp = box(6, 2, 4, MAT.concrete, -12, 1, -30)
    addClickable(wtp, 'Utilities', 'Water Treatment Plant', 'Treats raw river water (coagulation, flocculation, filtration, demineralization) for boiler feed and process use.', '#7c3aed')
    cyl(2, 2, 0.3, 24, MAT.water, -8, 0.15, -28)
    const sub = box(5, 2.5, 4, M(0x2a2a4a), 20, -0.25, -12)
    addClickable(sub, 'Utilities', 'Main Substation & Control Room', 'High-voltage substation and centralized DCS control room — manages all plant electrical and process control systems.', '#7c3aed')
    box(1.5, 2, 1.5, M(0x1a1a3a), 22, -0.5, -12)
    box(1.5, 2, 1.5, M(0x1a1a3a), 24, -0.5, -12)
  }

  // ZONE 8: RECREATION
  {
    box(16, 0.05, 12, M(0x2a5a1a), 36, 0.02, -24)
    const pool = box(10, 0.3, 5, MAT.pool, 36, 0.15, -25)
    addClickable(pool, 'Amenities', 'Recreation Pool', 'Staff recreational swimming pool and sports facility — part of mill housing complex for permanent employees.', '#7c3aed')
    for (let i = 0; i < 3; i++) box(0.3, 0.4, 5, M(0xf0f0e0), 31.2 + i * 5, 0.4, -25)
    const club = box(5, 2.5, 4, M(0xd0c8b0), 40, -0.25, -20)
    addClickable(club, 'Amenities', 'Staff Clubhouse & Canteen', 'Recreation center, canteen and meeting facilities for mill employees.', '#7c3aed')
    box(5, 0.5, 4, MAT.roofRed, 40, 1.75, -20)
    treeCluster(37, -26, 10, 8, 1.5, 3.5)
  }

  // ZONE 9: OFFICE COMPLEX
  {
    const office = box(5, 2.5, 5, M(0xc8b890), 10, 1.25, 14)
    addClickable(office, 'Administration', 'Administration & Main Gate', 'Main administration building, security post and visitor reception area at the main plant entrance.', '#7c3aed')
    box(5, 0.5, 5, M(0x3a3a3a), 10, 2.6, 14)
    box(5.1, 0.6, 0.3, M(0x1a5a9a), 10, 1.0, 11.6)
    box(8, 0.05, 5, M(0x2a2820), 10, 0.02, 20)
    for (let i = 0; i < 4; i++) box(0.1, 0.05, 4.5, MAT.white, 7 + i * 2.2, 0.04, 20)
  }

  // ZONE 10: COVERED CHIP STORAGE
  {
    const padE = box(14, 0.05, 8, MAT.sand, 27, 0.02, 11)
    addClickable(padE, 'Wood Handling', 'Covered Chip Storage', 'Secondary chip and bark storage area beside the pulp warehouse, sheltered under an open-sided canopy roof.', '#d97706')
    for (let col = 0; col < 3; col++) for (let row = 0; row < 2; row++) {
      const px = 21 + col * 6, pz = 8 + row * 6
      const pile = new THREE.Mesh(new THREE.SphereGeometry(2.6, 12, 8), (col + row) % 2 ? MAT.chip : MAT.bark)
      pile.scale.set(1.2, 0.4, 1.0)
      pile.position.set(px, 1.0, pz)
      pile.castShadow = true
      scene.add(pile)
    }
    for (let i = 0; i < 5; i++) cyl(0.12, 0.12, 3.2, 6, MAT.steel, 19 + i * 4, 3.2, 11)
    box(20, 0.3, 8, M(0x2a4a2a, 0x0a1a0a, 0.7), 27, 4.8, 11)
  }

  // ZONE 11: TANK FARM
  {
    for (let i = 0; i < 4; i++) {
      const t = cyl(1.3, 1.3, 5, 20, MAT.cream, -30 + i * 3.2, 2.5, -3)
      if (i === 1) addClickable(t, 'Utilities', 'Process & Chemical Storage Tanks', 'Bulk storage tanks for process chemicals and treated water supplying the mill.', '#7c3aed')
    }
    for (let i = 0; i < 3; i++) {
      cyl(0.85, 0.85, 3, 16, MAT.silver, -29 + i * 3.2, 1.5, -8)
      cyl(0.9, 0.9, 0.4, 16, M(0x2a7a4a), -29 + i * 3.2, 3.05, -8)
    }
  }

  // ZONE 12: CONVEYOR + WATER TANK
  {
    box(25, 0.4, 0.8, MAT.steel, 32, 3, 4)
    box(0.3, 4, 0.3, MAT.steel, 20, 2, 4)
    box(0.3, 4, 0.3, MAT.steel, 44, 2, 4)
    const conv = box(2, 2, 2, MAT.concrete, 46, 1, 4)
    addClickable(conv, 'Logistics', 'Jetty Conveyor Terminal', 'Belt conveyor system transports finished pulp bales from warehouse to river jetty for loading onto cargo ships for export.', '#7c3aed')
    const bigTank = cyl(2.2, 2.2, 6, 20, MAT.cream, 44, 3, 30)
    addClickable(bigTank, 'Utilities', 'Raw / Fire Water Storage Tank', 'Large-capacity water storage tank supplying fire protection and backup process water for the southeast yard area.', '#7c3aed')
  }

  // ── PROCESS FLOW ARROWS + ANIMATED PARTICLES ──
  {
    const S = 1.5 // spread factor
    const bluePipeMat = new THREE.MeshStandardMaterial({ color: 0x2563eb, emissive: 0x1d4ed8, emissiveIntensity: 0.4, transparent: true, opacity: 0.5 })
    const greenPipeMat = new THREE.MeshStandardMaterial({ color: 0x16a34a, emissive: 0x15803d, emissiveIntensity: 0.4, transparent: true, opacity: 0.5 })
    const orangePipeMat = new THREE.MeshStandardMaterial({ color: 0xea580c, emissive: 0xc2410c, emissiveIntensity: 0.3, transparent: true, opacity: 0.45 })

    function sp(p) { return [p[0] * S, p[1], p[2] * S] }

    function makeFlowTube(points, mat, radius = 0.25) {
      const scaled = points.map(p => new THREE.Vector3(p[0] * S, p[1], p[2] * S))
      const curve = new THREE.CatmullRomCurve3(scaled)
      const tubeGeo = new THREE.TubeGeometry(curve, 64, radius, 8, false)
      const tube = new THREE.Mesh(tubeGeo, mat)
      scene.add(tube)
      return curve
    }

    function addFlowParticles(curve, color, count = 8, speed = 0.15, size = 0.6) {
      for (let i = 0; i < count; i++) {
        const mat = new THREE.MeshBasicMaterial({ color, transparent: true, opacity: 0.9 })
        const mesh = new THREE.Mesh(new THREE.SphereGeometry(size, 8, 8), mat)
        scene.add(mesh)
        flowParticles.push({ mesh, curve, t: i / count, speed, size })
      }
    }

    // ── BLUE: Main Pulp Process Flow ──
    const c1 = makeFlowTube([[-30,1.2,-23],[-20,1.5,-22],[-10,1.5,-20],[-5,2,-22],[2,3.5,-14]], bluePipeMat)
    addFlowParticles(c1, 0x60a5fa, 6, 0.12)

    const c2 = makeFlowTube([[-2,3.5,-14],[2,5.5,-10],[2,3,-5],[8,2,-8]], bluePipeMat)
    addFlowParticles(c2, 0x60a5fa, 5, 0.14)

    const c3 = makeFlowTube([[8,2,-8],[14,4,-8],[14,2.5,-2],[14,4.5,5]], bluePipeMat)
    addFlowParticles(c3, 0x60a5fa, 5, 0.13)

    const c4 = makeFlowTube([[14,4.5,5],[0,3.5,4],[-4,2.5,4],[-12,2.5,4]], bluePipeMat)
    addFlowParticles(c4, 0x60a5fa, 6, 0.11)

    const c5 = makeFlowTube([[-4,2.5,4],[8,2.5,4],[18,2.5,4]], bluePipeMat)
    addFlowParticles(c5, 0x60a5fa, 4, 0.16)

    const c6 = makeFlowTube([[18,2.5,4],[32,3.5,4],[46,2.5,4]], bluePipeMat)
    addFlowParticles(c6, 0x60a5fa, 3, 0.1)

    // ── GREEN: White Liquor (Recausticizing → Digester) ──
    const c7 = makeFlowTube([[8,1,-22],[8,3,-18],[5,4,-14],[2,5.5,-10]], greenPipeMat, 0.2)
    addFlowParticles(c7, 0x4ade80, 5, 0.1, 0.5)

    // ── ORANGE: Black Liquor (Washing → Evaporator → Recovery Boiler) ──
    const c8 = makeFlowTube([[8,2,-8],[12,2.5,-10],[18,3.5,-14],[21,3.5,-14]], orangePipeMat, 0.2)
    addFlowParticles(c8, 0xfb923c, 5, 0.12, 0.5)

    const c9 = makeFlowTube([[21,3.5,-14],[15,4.5,-16],[8,5.5,-18]], orangePipeMat, 0.2)
    addFlowParticles(c9, 0xfb923c, 4, 0.11, 0.5)

    // ── GREEN: Smelt (Recovery Boiler → Recausticizing) ──
    const c10 = makeFlowTube([[8,0.5,-18],[8,0.5,-22]], greenPipeMat, 0.15)
    addFlowParticles(c10, 0x4ade80, 3, 0.2, 0.4)

    // ── ORANGE: Lime (Lime Kiln → Recausticizing) ──
    const c11 = makeFlowTube([[17,1.5,-22],[13,1.5,-22],[8,1,-22]], orangePipeMat, 0.15)
    addFlowParticles(c11, 0xfb923c, 3, 0.18, 0.4)

    // ── GREEN: Steam (Recovery Boiler → Power Boiler → Turbine) ──
    const c12 = makeFlowTube([[8,5.5,-18],[12,5.5,-18],[15,5.5,-18]], greenPipeMat, 0.15)
    addFlowParticles(c12, 0x4ade80, 3, 0.2, 0.4)

    const c13 = makeFlowTube([[15,4.5,-18],[15,2.5,-12]], greenPipeMat, 0.12)
    addFlowParticles(c13, 0x4ade80, 2, 0.25, 0.35)

    // ── Static Arrow Cones ──
    const arrowMat = new THREE.MeshBasicMaterial({ color: 0x60a5fa, transparent: true, opacity: 0.6 })
    ;[
      [-15*S,1.8,-21*S,-0.4],[-3*S,2.5,-16*S,0.3],[5*S,3.5,-9*S,0.5],
      [11*S,2.5,-6*S,0.2],[14*S,3,1*S,-0.1],[5*S,2.5,4*S,-0.3],
      [12*S,2.5,4*S,-0.2],[28*S,3.2,4*S,-0.1],
    ].forEach(([x,y,z,ry]) => {
      const cone = new THREE.Mesh(new THREE.ConeGeometry(0.4, 1, 6), arrowMat)
      cone.position.set(x, y, z)
      cone.rotation.y = ry
      cone.rotation.z = Math.PI / 2
      scene.add(cone)
    })

    const greenArrowMat = new THREE.MeshBasicMaterial({ color: 0x4ade80, transparent: true, opacity: 0.6 })
    ;[
      [8*S,1.5,-20*S,Math.PI/2+0.3],[5*S,3,-12*S,Math.PI/2+0.5],
    ].forEach(([x,y,z,ry]) => {
      const cone = new THREE.Mesh(new THREE.ConeGeometry(0.35, 0.9, 6), greenArrowMat)
      cone.position.set(x, y, z)
      cone.rotation.y = ry
      cone.rotation.z = Math.PI / 2
      scene.add(cone)
    })
  }

  // LABELS
  const labelsData = [
    { text: 'Log Storage Ponds', col: '#fbbf24', x: -30, y: 4, z: -24 },
    { text: 'Chip & Log Yard', col: '#fbbf24', x: -30, y: 4, z: 22 },
    { text: 'Recovery Boiler', col: '#f87171', x: 8, y: 12, z: -18 },
    { text: 'Digester', col: '#60a5fa', x: 2, y: 12, z: -10 },
    { text: 'Pulp Machine', col: '#60a5fa', x: -4, y: 5, z: 4 },
    { text: 'Warehouse', col: '#60a5fa', x: 18, y: 6, z: 4 },
    { text: 'Bleaching', col: '#34d399', x: 14, y: 6, z: -2 },
    { text: 'WWTP', col: '#2dd4bf', x: 30, y: 4, z: 24 },
    { text: 'Evaporators', col: '#f97316', x: 21, y: 6, z: -14 },
    { text: 'Lime Kiln', col: '#f97316', x: 17, y: 4, z: -22 },
    { text: 'HD Tower', col: '#38bdf8', x: 14, y: 10, z: 5 },
    { text: 'Pool', col: '#67e8f9', x: 36, y: 3, z: -25 },
    { text: 'Tank Farm', col: '#fbbf24', x: -25, y: 4, z: -3 },
    { text: 'Chip Storage', col: '#fbbf24', x: 27, y: 6, z: 11 }
  ]
  labelsData.forEach(l => {
    const s = makeLabel(l.text, l.col)
    s.position.set(l.x, l.y, l.z)
    scene.add(s)
  })

  updateCam()
}

function updateCam() {
  camera.position.set(
    orRadius * Math.sin(orPhi) * Math.sin(orTheta),
    orRadius * Math.cos(orPhi),
    orRadius * Math.sin(orPhi) * Math.cos(orTheta)
  )
  if (!currentLookAt) currentLookAt = targetLookAt.clone()
  currentLookAt.lerp(targetLookAt, 0.06)
  camera.lookAt(currentLookAt)
}

function toggleAutoRotate() {
  autoRotate.value = !autoRotate.value
}

function topView() {
  tPhi = 0.18; tRadius = 110; autoRotate.value = false
}

function resetView() {
  tTheta = -0.4; tPhi = 0.92; tRadius = 95; autoRotate.value = true
  activeStage.value = null
  stagePanelVisible.value = false
}

function goToStage(index) {
  const stage = stages[index]
  activeStage.value = index
  stagePanelVisible.value = false

  tTheta = stage.theta
  tPhi = stage.phi
  tRadius = stage.radius
  autoRotate.value = false

  targetLookAt = new THREE.Vector3(stage.lookAt[0], stage.lookAt[1], stage.lookAt[2])

  setTimeout(() => { stagePanelVisible.value = true }, 300)
}

let targetLookAt = new THREE.Vector3(0, 2, 2)

const raycaster = new THREE.Raycaster()
const mouse2 = new THREE.Vector2()

function onPointerDown(e) {
  isDrag = true
  prevX = e.clientX ?? e.touches?.[0]?.clientX ?? 0
  prevY = e.clientY ?? e.touches?.[0]?.clientY ?? 0
  dragStartX = prevX
  dragStartY = prevY
  autoRotate.value = false
}

function onPointerMove(e) {
  if (!isDrag) return
  const cx = e.clientX ?? e.touches?.[0]?.clientX ?? 0
  const cy = e.clientY ?? e.touches?.[0]?.clientY ?? 0
  tTheta -= (cx - prevX) * 0.006
  tPhi = Math.max(0.18, Math.min(1.42, tPhi + (cy - prevY) * 0.006))
  prevX = cx; prevY = cy
}

function onPointerUp() {
  isDrag = false
}

function onWheel(e) {
  tRadius = Math.max(20, Math.min(250, tRadius + e.deltaY * 0.08))
}

function onClick(e) {
  if (!renderer) return
  const dx = e.clientX - dragStartX
  const dy = e.clientY - dragStartY
  if (Math.abs(dx) > 5 || Math.abs(dy) > 5) return

  const rect = renderer.domElement.getBoundingClientRect()
  mouse2.x = ((e.clientX - rect.left) / rect.width) * 2 - 1
  mouse2.y = -((e.clientY - rect.top) / rect.height) * 2 + 1
  raycaster.setFromCamera(mouse2, camera)
  const hits = raycaster.intersectObjects(clickables, false)
  if (hits.length > 0 && hits[0].object.userData.name) {
    const u = hits[0].object.userData
    tooltip.x = Math.min(e.clientX - container.value.getBoundingClientRect().left + 18, container.value.clientWidth - 260)
    tooltip.y = Math.max(e.clientY - container.value.getBoundingClientRect().top - 70, 10)
    tooltip.visible = true
    tooltip.zone = u.zone
    tooltip.name = u.name
    tooltip.desc = u.desc
    tooltip.zoneColor = u.zoneColor || '#aaa'
  } else {
    tooltip.visible = false
  }
}

function onResize() {
  if (!container.value || !renderer || !camera) return
  const w = container.value.clientWidth
  const h = container.value.clientHeight
  camera.aspect = w / h
  camera.updateProjectionMatrix()
  renderer.setSize(w, h)
}

function animate(time) {
  animationId = requestAnimationFrame(animate)
  const dt = 0.016
  if (autoRotate.value) tTheta += 0.002
  orTheta += (tTheta - orTheta) * 0.06
  orPhi += (tPhi - orPhi) * 0.06
  orRadius += (tRadius - orRadius) * 0.06
  updateCam()

  // Animate flow particles
  for (const p of flowParticles) {
    p.t += p.speed * dt
    if (p.t > 1) p.t -= 1
    const pos = p.curve.getPoint(p.t)
    p.mesh.position.copy(pos)
    p.mesh.material.opacity = 0.6 + Math.sin(time * 0.005 + p.t * 20) * 0.3
  }

  renderer.render(scene, camera)
}

onMounted(() => {
  nextTick(() => {
    buildScene()
    animate(0)
    window.addEventListener('resize', onResize)
    renderer.domElement.addEventListener('pointerdown', onPointerDown)
    renderer.domElement.addEventListener('pointermove', onPointerMove)
    renderer.domElement.addEventListener('pointerup', onPointerUp)
    renderer.domElement.addEventListener('wheel', onWheel, { passive: true })
    renderer.domElement.addEventListener('click', onClick)
  })
})

onUnmounted(() => {
  window.removeEventListener('resize', onResize)
  if (animationId) cancelAnimationFrame(animationId)
  if (renderer) {
    renderer.domElement.removeEventListener('pointerdown', onPointerDown)
    renderer.domElement.removeEventListener('pointermove', onPointerMove)
    renderer.domElement.removeEventListener('pointerup', onPointerUp)
    renderer.domElement.removeEventListener('wheel', onWheel)
    renderer.domElement.removeEventListener('click', onClick)
    renderer.dispose()
    container.value?.removeChild(renderer.domElement)
  }
  clickables.length = 0
  flowParticles.length = 0
  scene?.clear()
})
</script>
