terraform {
    # определяет требуемых провайдеров
    required_providers {
        twc = {
            source = "tf.timeweb.cloud/timeweb-cloud/timeweb-cloud"
        }
    }
    required_version = ">= 0.13"
}

provider "twc" {    
    token = "..."
}

data "twc_configurator" "configurator" {
    location = "ru-1"
    disk_type = "nvme"
}

data "twc_os" "os" {
    name = "ubuntu"
    version = "22.04"
}

resource "twc_server" "example-server" {
    name = "Example server"
    os_id = data.twc_os.os.id

    configuration {
        configurator_id = data.twc_configurator.configurator.id
        disk = 1024 * 15
        cpu = 1
        ram = 1024
    }
}