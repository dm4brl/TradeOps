provider "google" {
  project = "your-project-id"
  region  = "us-central1"
}

resource "google_storage_bucket" "my_bucket" {
  name          = "my-bucket"
  location      = "US"
  force_destroy = true
}

resource "google_cloudfunctions_function" "supplier_service" {
  name        = "supplier-service"
  runtime     = "go113"
  source_archive_bucket = google_storage_bucket.my_bucket.name
  source_archive_object = "supplier-service.zip"
  entry_point = "SupplierService"
  trigger_http = true
  available_memory_mb = 256
  project     = "your-project-id"
  region      = "us-central1"
}

// Define other functions and resources here...
