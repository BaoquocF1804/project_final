# Project_final

Project use Golang để xây dựng 

Các framework sử dụng:
+ gRPC và gRPC gateway
+ Sqlc 
+ Mockery 

Mô tả: 

- Project hiện thực 2 gRPC sever và 1 API_http:
    + gRPC1_user : có phương thức như CreateUser, GetUser
    + gRPC2_account: có các phương thức như GetAccount, CreateAccount
    + API_http : nhận đầu vào thông tin ID Account sau đó xuất ra thông tin user bằng cách lấy API client của gate_account sau đó xuất ra "Owner". Từ "Owner" lấy thông tin gọi đến gate_user để trả ra thông tin User và xuất ra API http bằng file json.
  
