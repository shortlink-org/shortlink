/**
 * Runtime проверка разрешений Permissions API
 * 
 * Используется для валидации и логирования разрешений в runtime.
 * Полезно для отладки и аудита безопасности.
 */

/**
 * Проверяет доступность Permissions API
 * 
 * @returns true если Permissions API доступен, иначе false
 */
export function isPermissionsAPIAvailable(): boolean {
  return typeof process.permission !== "undefined";
}

/**
 * Проверяет наличие конкретного разрешения
 * 
 * @param resource - тип ресурса ('fs.read', 'fs.write', 'env', 'net')
 * @param name - имя ресурса (путь для fs, имя переменной для env, хост для net)
 * @returns true если разрешение есть, иначе false
 */
export function hasPermission(
  resource: "fs.read" | "fs.write" | "env" | "net",
  name: string
): boolean {
  if (!isPermissionsAPIAvailable()) {
    // Если Permissions API не доступен, считаем что разрешение есть
    // (для обратной совместимости)
    return true;
  }

  try {
    return process.permission?.has(resource, name) ?? false;
  } catch (error) {
    console.warn(`[Permissions] Failed to check permission for ${resource}:${name}`, error);
    return false;
  }
}

/**
 * Проверяет все необходимые разрешения для Proxy Service
 * 
 * @returns объект с результатами проверки каждого разрешения
 */
export function validateRequiredPermissions(): {
  fsRead: boolean;
  env: boolean;
  net: boolean;
  allGranted: boolean;
} {
  // Проверяем fs.read для рабочей директории (работает и локально, и в Docker)
  const cwd = process.cwd();
  const fsRead = hasPermission("fs.read", cwd) || hasPermission("fs.read", "/app");
  
  // env проверка не работает через флаги в Node.js 25, контролируется на уровне контейнера
  // В Node.js 25 нет детального контроля env через флаги, только через контейнер
  const env = true; // Переменные окружения контролируются на уровне Kubernetes/Docker
  
  // net проверка - в Node.js 25 --allow-net разрешает всю сеть (экспериментально)
  const net = hasPermission("net", "*") || hasPermission("net", "localhost:3020");

  return {
    fsRead,
    env,
    net,
    allGranted: fsRead && env && net,
  };
}

/**
 * Логирует текущие разрешения (для отладки)
 */
export function logPermissions(): void {
  if (!isPermissionsAPIAvailable()) {
    console.log("[Permissions] Permissions API not available (running without restrictions)");
    return;
  }

  console.log("[Permissions] Permissions API enabled");
  
  const permissions = validateRequiredPermissions();
  console.log("[Permissions] Required permissions:", {
    "fs.read": permissions.fsRead ? "✓" : "✗",
    "env": "✓ (controlled by container)", // В Node.js 25 env контролируется на уровне контейнера
    "net": permissions.net ? "✓" : "✗",
  });

  // Проверяем только критичные разрешения (fs.read и net)
  // env всегда true, так как контролируется на уровне контейнера
  if (!permissions.fsRead || !permissions.net) {
    console.warn("[Permissions] ⚠️  Some required permissions are missing!");
  } else {
    console.log("[Permissions] ✓ All required permissions granted");
  }
}

