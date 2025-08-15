cat > setup_env.sh << 'EOF'
#!/bin/bash

export ENV=production
echo "Setting up $ENV environment"

calculate() {
    echo $(($1 + $2))
}
EOF

# ------------------------------------------------------------------------

source setup_env.sh
calculate 5 5