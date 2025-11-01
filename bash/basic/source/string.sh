eval "$(cat <<'EOF'
#!/bin/bash

export ENV=production
echo "Setting up $ENV environment"

calculate() {
    echo $(($1 + $2))
}
EOF
)"

calculate 5 5

echo $ENV