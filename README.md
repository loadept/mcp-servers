# Servidores MCP (Model Context Protocol)

Este repositorio contiene servidores MCP para diferentes propósitos, incluyendo la interacción con el sistema de archivos y bases de datos PostgreSQL. Los servidores MCP permiten que los asistentes de IA interactúen de forma segura con recursos externos utilizando herramientas especializadas.

## 📋 Descripción General

Los servidores MCP (Model Context Protocol) actúan como puentes entre los asistentes de IA y recursos externos, proporcionando herramientas estructuradas y seguras para realizar operaciones específicas. Este repositorio incluye dos servidores principales:

### 🗂️ Servidor de Sistema de Archivos (`filesystem/`)
Un servidor MCP escrito en Python que permite la interacción segura con el sistema de archivos local.

### 🐘 Servidor PostgreSQL (`pg/`)
Un servidor MCP escrito en Go que facilita la consulta e inspección de bases de datos PostgreSQL.

## 🚀 Servidores Disponibles

### 1. Servidor de Sistema de Archivos

**Tecnología:** Python 3.13+  
**Paquete:** `loadept-mcp-filesystem`  
**Versión:** 0.3.3

#### Características:
- ✅ Navegación segura de directorios (con validación de rutas)
- ✅ Listado de contenidos de directorios
- ✅ Lectura de archivos de texto
- ✅ Escritura de archivos
- ✅ Búsqueda de archivos por nombre o patrón
- ✅ Creación de directorios
- ✅ Apertura de archivos con aplicación predeterminada
- ✅ Soporte multiplataforma (Windows, macOS, Linux)

#### Herramientas Disponibles:
- `list_directory`: Lista el contenido de un directorio
- `find_results`: Busca archivos que coincidan con un patrón específico
- `read_content`: Lee el contenido de un archivo
- `write_content`: Escribe contenido en un archivo
- `open_file`: Abre un archivo con la aplicación predeterminada
- `create_directory`: Crea un nuevo directorio

#### Seguridad:
- Validación de rutas para prevenir acceso fuera del directorio base
- Bloqueo de rutas absolutas y navegación hacia arriba (`../`)
- Configuración de directorio base mediante variable de entorno `BASE_PATH`

### 2. Servidor PostgreSQL

**Tecnología:** Go 1.25+  
**Versión:** 0.1.1

#### Características:
- ✅ Ejecución segura de consultas SELECT (solo lectura)
- ✅ Inspección de estructura de tablas
- ✅ Listado de tablas por esquema
- ✅ Información detallada de columnas
- ✅ Paginación de resultados
- ✅ Transacciones de solo lectura
- ✅ Límite de 50 filas por consulta para optimización

#### Herramientas Disponibles:
- `execute_query`: Ejecuta consultas SQL de solo lectura en PostgreSQL
- `get_table_info`: Obtiene información detallada sobre una tabla específica
- `list_tables`: Lista todas las tablas disponibles en un esquema

#### Seguridad:
- Solo permite consultas SELECT (sin modificaciones de datos)
- Transacciones de solo lectura con nivel de aislamiento Read Committed
- Límite automático de resultados para prevenir sobrecarga

## 🛠️ Instalación y Configuración

### Prerrequisitos

#### Para el Servidor de Sistema de Archivos:
```bash
# Verificar instalación de UV (gestor de paquetes Python)
uv --version
```

Si no tienes UV instalado:
```bash
# En macOS y Linux
curl -LsSf https://astral.sh/uv/install.sh | sh

# En Windows
powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex"
```

#### Para el Servidor PostgreSQL:
- Go 1.25+
- Acceso a una base de datos PostgreSQL

### Configuración en Clientes MCP

#### 📝 VS Code
Archivo: `$env:USERPROFILE\AppData\Roaming\Code\User\mcp.json`
```json
{
  "servers": {
    "loadept-mcp-filesystem": {
      "type": "stdio",
      "command": "uvx",
      "args": ["loadept-mcp-filesystem"],
      "env": {
        "BASE_PATH": "C:\\Ruta\\Al\\Directorio\\Base"
      }
    },
    "postgres-mcp": {
      "type": "stdio",
      "command": "path/to/pg-mcp",
      "env": {
        "POSTGRES_URI": "postgres://usuario:contraseña@localhost:5432/basedatos?sslmode=disable"
      }
    }
  }
}
```

#### 🖱️ Cursor AI
Archivo: `$env:USERPROFILE\.cursor\mcp.json`
```json
{
  "mcpServers": {
    "loadept-mcp-filesystem": {
      "command": "uvx",
      "args": ["loadept-mcp-filesystem"],
      "env": {
        "BASE_PATH": "C:\\Ruta\\Al\\Directorio\\Base"
      }
    },
    "postgres-mcp": {
      "command": "path/to/pg-mcp",
      "env": {
        "POSTGRES_URI": "postgres://usuario:contraseña@localhost:5432/basedatos?sslmode=disable"
      }
    }
  }
}
```

#### 🤖 Claude Desktop
Archivo: `$env:USERPROFILE\AppData\Roaming\Claude\claude_desktop_config.json`
```json
{
  "mcpServers": {
    "loadept-mcp-filesystem": {
      "command": "uvx",
      "args": ["loadept-mcp-filesystem"],
      "env": {
        "BASE_PATH": "C:\\Ruta\\Al\\Directorio\\Base"
      }
    },
    "postgres-mcp": {
      "command": "path/to/pg-mcp",
      "env": {
        "POSTGRES_URI": "postgres://usuario:contraseña@localhost:5432/basedatos?sslmode=disable"
      }
    }
  }
}
```

## 🔧 Desarrollo

### Estructura del Proyecto

```
mcp-servers/
├── filesystem/                 # Servidor de sistema de archivos (Python)
│   ├── loadept_mcp_filesystem/
│   │   ├── schemas/           # Esquemas de herramientas Pydantic
│   │   ├── tools/            # Implementación de herramientas
│   │   ├── utils/            # Utilidades y validadores
│   │   └── server.py         # Servidor MCP principal
│   ├── pyproject.toml        # Configuración del proyecto Python
│   └── README.md
│
├── pg/                        # Servidor PostgreSQL (Go)
│   ├── cmd/
│   │   ├── mcp/main.go       # Punto de entrada del servidor MCP
│   │   └── db/main.go        # Utilidad de conexión a DB
│   ├── internal/
│   │   ├── config/           # Configuración y variables de entorno
│   │   ├── di/               # Inyección de dependencias
│   │   ├── domain/           # Modelos de dominio
│   │   ├── infra/            # Infraestructura (conexión DB)
│   │   ├── repository/       # Capa de acceso a datos
│   │   ├── service/          # Lógica de negocio
│   │   └── transport/        # Transporte MCP
│   └── go.mod
│
└── .github/workflows/         # CI/CD con GitHub Actions
```

### Construcción desde el Código Fuente

#### Servidor de Sistema de Archivos:
```bash
cd filesystem
uv venv
source .venv/bin/activate  # En Windows: .venv\Scripts\activate
uv pip install -e .
```

#### Servidor PostgreSQL:
```bash
cd pg
go mod download
go build -o pg-mcp cmd/mcp/main.go
```

## 🔐 Variables de Entorno

### Servidor de Sistema de Archivos:
- `BASE_PATH`: Directorio base para operaciones de archivos (obligatorio para seguridad)

### Servidor PostgreSQL:
- `POSTGRES_URI`: URI de conexión a PostgreSQL (formato: `postgres://usuario:contraseña@host:puerto/basedatos`)

## 🏗️ Arquitectura

### Servidor de Sistema de Archivos:
- **Patrón Registro**: Decoradores para registro automático de herramientas y esquemas
- **Validación de Seguridad**: Validación estricta de rutas para prevenir vulnerabilidades
- **Manejo de Errores**: Gestión robusta de errores con mensajes descriptivos

### Servidor PostgreSQL:
- **Arquitectura Hexagonal**: Separación clara entre dominio, infraestructura y transporte
- **Inyección de Dependencias**: Gestión centralizada de dependencias
- **Patrón Repository**: Abstracción del acceso a datos
- **Seguridad por Defecto**: Solo operaciones de lectura con transacciones seguras
