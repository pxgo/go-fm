// 最简单的Service Worker，仅用于PWA安装

// 安装事件 - 立即激活
self.addEventListener('install', function(event) {
  self.skipWaiting();
});

// 激活事件 - 立即获取控制权
self.addEventListener('activate', function(event) {
  event.waitUntil(self.clients.claim());
});
