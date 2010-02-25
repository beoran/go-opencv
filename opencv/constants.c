/*
#include <opencv/cv.h>




#include <opencv/highgui.h>
*/
#include <opencv/cvaux.h>
/*
typedef struct ImgObsInfo $CvImgObsInfo; 
typedef struct EHMMState $CvEHMMState;
typedef struct EHMM $CvEHMM;
typedef struct GraphWeightedVtx $CvGraphWeightedVtx;
typedef struct GraphWeightedEdge $CvGraphWeightedEdge;
typedef struct CliqueFinder $CvCliqueFinder;
typedef struct StereoLineCoeff $CvStereoLineCoeff;
typedef struct Camera $CvCamera;
typedef struct StereoCamera $CvStereoCamera;
typedef struct ContourOrientation $CvContourOrientation;
typedef struct GLCM $CvGLCM
typedef struct FaceTracker $CvFaceTracker
typedef struct Face $CvFace;
typedef struct VoronoiSite2D $CvVoronoiSite2D;
typedef struct VoronoiEdge2D $CvVoronoiEdge2D;
typedef struct VoronoiNode2D $CvVoronoiNode2D;
typedef struct VoronoiDiagram2D $CvVoronoiDiagram2D;
typedef struct LCMEdge $CvLCMEdge;
typedef struct LCMNode $CvLCMNode;
typedef struct BGStatModel $CvBGStatModel;
typedef struct FGDStatModelParams $CvFGDStatModelParams;
typedef struct BGPixelCStatTable $CvBGPixelCStatTable;
typedef struct BGPixelCCStatTable $CvBGPixelCCStatTable;
typedef struct BGPixelStat $CvBGPixelStat;
typedef struct FGDStatModel $CvFGDStatModel;
typedef struct GaussBGStatModelParams $CvGaussBGStatModelParams;
typedef struct GaussBGValues $CvGaussBGValues;
typedef struct GaussBGPoint $CvGaussBGPoint;
typedef struct GaussBGModel $CvGaussBGModel;
typedef struct BGCodeBookElem $CvBGCodeBookElem;
typedef struct BGCodeBookModel $CvBGCodeBookModel;
typedef struct RandState $CvRandState;

typedef struct FeatureTree $CvFeatureTree;
typedef struct LSH $CvLSH;
typedef struct SHOperations $CvLSHOperations;

typedef struct SURFPoint $CvSURFPoint;
typedef struct SURFParams $CvSURFParams;
typedef struct MSERParams $CvMSERParams;
typedef struct StarKeypoint $CvStarKeypoint;
typedef struct StarDetectorParams $CvStarDetectorParams;
typedef struct POSITObject $CvPOSITObject; CvPOSITObject;

typedef struct StereoBMState $CvStereoBMState;
typedef struct StereoGCState $CvStereoGCState;
typedef struct Moments $CvMoments;
typedef struct HuMoments $CvHuMoments;
typedef struct ConnectedComp $CvConnectedComp;

typedef struct ContourScanner $_CvContourScanner;
typedef struct ChainPtReader $CvChainPtReader;
typedef struct ContourTree $CvContourTree;
typedef struct ConvexityDefect $CvConvexityDefect;
typedef struct QuadEdge2D $CvQuadEdge2D;
typedef struct Subdiv2DPoint $CvSubdiv2DPoint;
typedef struct Subdiv2D $CvSubdiv2D;
typedef struct Matrix3 $CvMatrix3;
typedef struct ConDensation $CvConDensation;
typedef struct Kalman $CvKalman;

typedef struct HaarFeature $CvHaarFeature;
typedef struct HaarClassifier $CvHaarClassifier;
typedef struct HaarStageClassifier $CvHaarStageClassifier;
typedef struct HidHaarClassifierCascade $CvHidHaarClassifierCascade; 
typedef struct HaarClassifierCascade $CvHaarClassifierCascade;
typedef struct AvgComp $CvAvgComp;        
typedef struct NArrayIterator $CvNArrayIterator;
typedef struct GraphScanner $CvGraphScanner;

typedef struct Font $CvFont;
typedef struct TreeNodeIterator $CvTreeNodeIterator;

typedef struct FuncTable $CvFuncTable;
typedef struct BigFuncTable $CvBigFuncTable;
typedef struct Image        $_IplImage;
typedef struct TileInfo     $_IplTileInfo;
typedef struct ROI          $_IplROI;
typedef struct $ConvKernel  $_IplConvKernel;
typedef struct $ConvKernelFP $_IplConvKernelFP;
typedef struct Mat $CvMat;
typedef struct MatND $CvMatND;
typedef struct Set $CvSet;
typedef struct SparseMat $CvSparseMat;
typedef struct SparseNode $CvSparseNode;
typedef struct SparseMatIterator $CvSparseMatIterator;
typedef struct Histogram $CvHistogram;
typedef struct Rect $CvRect;
typedef struct TermCriteria $CvTermCriteria;
typedef struct Point $CvPoint;
typedef struct Point2D32f $CvPoint2D32f;
typedef struct Point3D32f $CvPoint3D32f;
typedef struct Point2D64f $CvPoint2D64f;
typedef struct Point3D64f $CvPoint3D64f;
typedef struct Size2D32f $CvSize2D32f;
typedef struct Box2D $CvBox2D;
typedef struct LineIterator $CvLineIterator;
typedef struct Slice $CvSlice;
typedef struct Scalar $CvScalar;
typedef struct MemBlock $CvMemBlock;
typedef struct MemStorage $CvMemStorage;
typedef struct MemStoragePos $CvMemStoragePos;
typedef struct SeqBlock $CvSeqBlock;
typedef struct Seq $CvSeq;
typedef struct SetElem $CvSetElem;
typedef struct Set $CvSet;
typedef struct GraphEdge $CvGraphEdge;
typedef struct GraphVtx $CvGraphVtx;
typedef struct GraphVtx2D $CvGraphVtx2D;
typedef struct Graph $CvGraph;
typedef struct Chain $CvChain;
typedef struct Contour $CvContour;
typedef struct SeqWriter $CvSeqWriter;
typedef struct SeqReader $CvSeqReader;
typedef struct FileStorage $CvFileStorage; 
typedef struct AttrList $CvAttrList;
typedef struct TypeInfo $CvTypeInfo;
typedef struct String $CvString;
typedef struct StringHashNode $CvStringHashNode;
typedef struct GenericHash $CvGenericHash; 
typedef struct FileNodeHash CvFileNodeHash;
typedef struct FileNode $CvFileNode;
typedef struct PluginFuncInfo $CvPluginFuncInfo;
typedef struct ModuleInfo $CvModuleInfo;
typedef struct Capture $CvCapture; CvCapture;
typedef struct VideoWriter $CvVideoWriter; 
typedef struct ParamLattice $CvParamLattice;
typedef struct CNNLayer $CvCNNLayer; 
typedef struct CNNetwork $CvCNNetwork;
typedef struct CNNLayer $CvCNNLayer;
typedef struct CNNConvolutionLayer $CvCNNConvolutionLayer;
typedef struct CNNSubSamplingLayer $CvCNNSubSamplingLayer;
typedef struct CNNFullConnectLayer $CvCNNFullConnectLayer;
typedef struct CNNetwork $CvCNNetwork;
typedef struct CNNStatModel $CvCNNStatModel;
typedef struct CNNStatModelParams $CvCNNStatModelParams;
typedef struct CrossValidationParams $CvCrossValidationParams;
typedef struct CrossValidationModel $CvCrossValidationModel;


enum
{
        $EVENT_MOUSEMOVE = CV_EVENT_MOUSEMOVE,
        $EVENT_LBUTTONDOWN = CV_EVENT_LBUTTONDOWN,
        $EVENT_RBUTTONDOWN = CV_EVENT_RBUTTONDOWN,
        $EVENT_MBUTTONDOWN = CV_EVENT_MBUTTONDOWN,
        $EVENT_LBUTTONUP = CV_EVENT_LBUTTONUP,
        $EVENT_RBUTTONUP = CV_EVENT_RBUTTONUP,
        $EVENT_MBUTTONUP = CV_EVENT_MBUTTONUP,
        $EVENT_LBUTTONDBLCLK = CV_EVENT_LBUTTONDBLCLK,
        $EVENT_RBUTTONDBLCLK = CV_EVENT_RBUTTONDBLCLK,
        $EVENT_MBUTTONDBLCLK = CV_EVENT_MBUTTONDBLCLK,
        $EVENT_FLAG_LBUTTON = CV_EVENT_FLAG_LBUTTON,
        $EVENT_FLAG_RBUTTON = CV_EVENT_FLAG_RBUTTON,
        $EVENT_FLAG_MBUTTON = CV_EVENT_FLAG_MBUTTON,
        $EVENT_FLAG_CTRLKEY = CV_EVENT_FLAG_CTRLKEY,
        $EVENT_FLAG_SHIFTKEY = CV_EVENT_FLAG_SHIFTKEY,
        $EVENT_FLAG_ALTKEY = CV_EVENT_FLAG_ALTKEY,
        $LOAD_IMAGE_UNCHANGED = CV_LOAD_IMAGE_UNCHANGED,
        $LOAD_IMAGE_GRAYSCALE = CV_LOAD_IMAGE_GRAYSCALE,
        $LOAD_IMAGE_COLOR = CV_LOAD_IMAGE_COLOR,
        $LOAD_IMAGE_ANYDEPTH = CV_LOAD_IMAGE_ANYDEPTH,
        $LOAD_IMAGE_ANYCOLOR = CV_LOAD_IMAGE_ANYCOLOR,
        $IMWRITE_JPEG_QUALITY = CV_IMWRITE_JPEG_QUALITY,
        $IMWRITE_PNG_COMPRESSION = CV_IMWRITE_PNG_COMPRESSION,
        $IMWRITE_PXM_BINARY = CV_IMWRITE_PXM_BINARY,
};

*/
