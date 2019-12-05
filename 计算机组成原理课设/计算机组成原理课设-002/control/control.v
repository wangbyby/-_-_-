module Controll(
    input [31:0] order,
    input clk,
    output reg [2:0] sel_ALU,
    output  reg  beq,Extop,j,jal,jr,lui,rd,imm_to_ALU,GPR_write, RAM_write, RAM_to_GPR
);
// beq 指令
// Extop 符号拓展
//j 指令
//jr 指令
// lui 指令
wire [5:0] op ,order_func;
assign op = order[31:26];
assign order_func = order[5:0];
always @(order or posedge clk) begin
    case(op)
        6'b000000: begin // r型 指令
            rd <= 1'b1;
            GPR_write <= 1'b1;
            if (order_func == 'b101010) begin // slt
                sel_ALU <= 3'b011;
                end
            else if (order_func == 'b100001) begin // addu
                sel_ALU <= 3'b000;
                end
            else if(order_func == 'b100011) begin // subu
                sel_ALU <= 3'b001;
                end
            else if (order_func == 'b001000) begin //jr
                jr <= 1'b1;
                end
            else begin // 其他
                sel_ALU <= 3'b000;
                end
            end
        6'b001000: //addi
            begin
            rd <= 1'b0;
            sel_ALU <= 3'b100;
            imm_to_ALU <= 1'b1;
            GPR_write <= 1'b1;
            end
        6'b001001:    //addiu
            begin
              rd <= 1'b0;
            sel_ALU <= 3'b000;
            imm_to_ALU <= 1'b1;
            GPR_write <= 1'b1;
            end
        6'b000100:    //beq
        begin
            sel_ALU <= 3'b001;
            beq <= 1'b1;
            end
        6'b001111:    //lui
        begin
            rd <= 1'b0;
            sel_ALU <= 3'b101;
            lui <= 1'b1;
            imm_to_ALU <= 1'b1;
            GPR_write <= 1'b1;
            end
        6'b100011:    //lw
          begin
            rd <= 1'b0;
            sel_ALU <= 3'b000;
            imm_to_ALU <= 1'b1;
            RAM_to_GPR <= 1'b1;
            GPR_write <= 1'b1;
            Extop <= 1'b1;
            end
        6'b001101:    //ori
          begin
            sel_ALU <= 3'b010;
            GPR_write <= 1'b1;
            imm_to_ALU <= 1'b1;
            rd <= 1'b0;
            end
        6'b101011:    //sw
          begin
            sel_ALU <= 3'b000;
            imm_to_ALU <= 1'b1;
            RAM_write <= 1'b1;
            Extop <= 1'b1;
            rd <= 1'b0;
            end
        6'b000010:
            //j
            begin
            j <= 1'b1;
            rd <= 1'b0;
            end
        6'b000011:    //jal
            begin 
            rd <= 1'b0;
            jal <= 1'b1;
          end
    default: begin
            j <= 0'b0;
          rd <= 1'b0;
          end
    endcase
end

endmodule // Controll