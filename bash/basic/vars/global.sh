SC_COMMON_LOGGING_CURRENT_SECTION_NAME=""
sc::common::logging::current_section_name() {
    echo ${SC_COMMON_LOGGING_CURRENT_SECTION_NAME}
}

# ------------------------------------------------------------------------

sc::common::logging::section_start() {
    local section_title="${1}"
    SC_COMMON_LOGGING_CURRENT_SECTION_NAME=${section_title}
}

sc::common::logging::section_end() {
  local section_title="${1}"
  #...
  SC_COMMON_LOGGING_CURRENT_SECTION_NAME=""
}

# visual test
# ------------------------------------------------------------------------

echo "cur_section: $(sc::common::logging::current_section_name)"

sc::common::logging::section_start "this_section"
echo "cur_section: $(sc::common::logging::current_section_name)"
sc::common::logging::section_end "this_section"

echo "cur_section: $(sc::common::logging::current_section_name)"

sc::common::logging::section_start "other_section"
echo "cur_section: $(sc::common::logging::current_section_name)"
sc::common::logging::section_end "other_section"

# ------------------------------------------------------------------------