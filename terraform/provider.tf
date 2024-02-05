provider "google" {
  credentials = file("/path/to/your/credentials.json")
  project     = var.project_id
  region      = var.region
}
