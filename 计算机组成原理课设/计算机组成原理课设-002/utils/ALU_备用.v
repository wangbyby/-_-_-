module ALU(clk,inA,inB,imm16,aluSelect,outC,flag_overflow_pos,flag_zero);

input [31:0]inA,inB; //数据线
input [15:0] imm16;
//选择
input [2:0]aluSelect;
input clk; // 时钟
//输出
output [31:0] outC;
// 溢出,0
output flag_overflow_pos,flag_zero; 

// 
reg [31:0] outC;
reg flag_overflow_pos,flag_zero;
reg [32:0] tmp;
always @(posedge clk or aluSelect or inA or inB ) 
    begin
    case (aluSelect)
        3'b000: outC <= $unsigned(inA)+$unsigned(inB); 
                //unsign  addu  addiu lw sw
        3'b001: begin 
                    outC <= $unsigned(inA) - $unsigned(inB); //unsign subu
                    flag_zero <= &(~outC); // beq 信号
                end
        3'b010: outC <= inA|inB; //ori
        3'b011: if (inA < inB) begin //slt
                outC <= 1;
            end else begin
                outC <= 0;
            end
        3'b100:  begin// addi 溢出
                    tmp <= {1'b0,inA} + {1'b0,inB};
                    outC <= tmp[31:0];
                    flag_overflow_pos <= tmp[32]^tmp[31]; //溢出判断
                    // if (tmp[32:31]==2'b00 || tmp[32:31]==2'b11 ) begin
                    //     flag_overflow_pos <= 1'b1;
                    // end else begin
                    //     flag_overflow_pos <= 1'b0;
                    // end
                    //flag_overflow_pos <= inA[31]^inB[31]^outC[31]^tmp[31];
                end
        3'b101: // lui
            outC <={imm16,16{1'b0}}; // lui
        default: outC <= 32'b0;
    endcase
    end
endmodule //in/ inALU A,B,aluSelect,outC