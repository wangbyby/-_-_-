module ALU(clk,inA,inB,aluSelect,outC,flag_overflow_pos,flag_zero);

input [31:0]inA,inB; //数据线
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
                end
        3'b101:
            outC <= inB; // lui
        default: outC <= inA+inB;
    endcase
    end
endmodule //in/ inALU A,B,aluSelect,outC