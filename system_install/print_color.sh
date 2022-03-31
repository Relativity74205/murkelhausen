function print_color(){
  NC='\033[0m' # No Color

  case $1 in
    "green") COLOR='\033[0;32m' ;;
    "red") COLOR='\033[0;31m' ;;
    "cyan") COLOR='\033[0;36m' ;;
    "purple") COLOR='\033[0;35m' ;;
    "yellow") COLOR='\033[1;33m' ;;
    "*") COLOR='\033[0m' ;;
  esac

  echo -e "${COLOR} $2 ${NC}"
}
